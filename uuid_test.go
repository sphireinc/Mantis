package mantis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

type marshalTestJSON struct {
	Pointer1 *UUID `json:"ptr1"`
	Pointer2 *UUID `json:"ptr2,omitempty"`
	Value    UUID  `json:"value"`
}

func TestMarshalUnMarshalJSON(t *testing.T) {
	u1 := MustParseUUIDString("5ebd21f5-73bd-4574-9598-68f11584e266")
	u2 := MustParseUUIDString("{5161487f-c712-4689-a12a-b391ab7eb423}")
	u3 := MustParseUUIDString("urn:uuid:02ed992a-4082-4981-af49-e4423d3e13b8")
	tests := []struct {
		v *marshalTestJSON
		e []byte
	}{
		{v: &marshalTestJSON{Pointer1: &u1, Pointer2: &u2, Value: u3}, e: []byte(`{"ptr1":"5ebd21f5-73bd-4574-9598-68f11584e266","ptr2":"5161487f-c712-4689-a12a-b391ab7eb423","value":"02ed992a-4082-4981-af49-e4423d3e13b8"}`)},
		{v: &marshalTestJSON{Pointer1: &u1, Pointer2: nil, Value: u3}, e: []byte(`{"ptr1":"5ebd21f5-73bd-4574-9598-68f11584e266","value":"02ed992a-4082-4981-af49-e4423d3e13b8"}`)},
		{v: &marshalTestJSON{Pointer1: nil, Pointer2: &u2, Value: u3}, e: []byte(`{"ptr1":null,"ptr2":"5161487f-c712-4689-a12a-b391ab7eb423","value":"02ed992a-4082-4981-af49-e4423d3e13b8"}`)},
		{v: &marshalTestJSON{Pointer1: &u1, Pointer2: &u2}, e: []byte(`{"ptr1":"5ebd21f5-73bd-4574-9598-68f11584e266","ptr2":"5161487f-c712-4689-a12a-b391ab7eb423","value":"00000000-0000-0000-0000-000000000000"}`)},
		{v: &marshalTestJSON{}, e: []byte(`{"ptr1":null,"value":"00000000-0000-0000-0000-000000000000"}`)},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			bs, err := json.Marshal(test.v)
			if err != nil {
				t.Fatal("unexpected marshal error")
			}
			if !bytes.Equal(test.e, bs) {
				t.Fatalf("unexpected json body. Expected\n%s\nActual\n%s", string(test.e), string(bs))
			}
			var v marshalTestJSON
			if err = json.Unmarshal(bs, &v); err != nil {
				t.Fatal("unexpected unmarshal error")
			}
			if v.Value != test.v.Value {
				t.Fatalf("unexpected unmarshal Value. Expected\n%v\nActual\n%v", test.v.Value, v.Value)
			}
			if !v.Pointer1.Equals(test.v.Pointer1) {
				t.Fatalf("unexpected unmarshal Pointer1. Expected\n%v\nActual\n%v", test.v.Pointer1, v.Pointer1)
			}
			if !v.Pointer2.Equals(test.v.Pointer2) {
				t.Fatalf("unexpected unmarshal Pointer2. Expected\n%v\nActual\n%v", test.v.Pointer2, v.Pointer2)
			}
			//t.Logf("%s", string(bs))
		})
	}
}
func TestUnmarshalUUIDError(t *testing.T) {
	testStrings := []string{
		"",                   // empty
		"e8c8cec324e9445aa0", // too short
		"e8c8cec324e9445aa086f021ecbac4ddaaaaaaaa",          // too long
		"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",                  // not hex
		"e8c8cec3-24e944-5a-a086f021ecbac-4dd",              // dashes misplaced
		"{xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx}",            // not hex (braced)
		"{e8c8cec3-24e9-445a-a086-f021ecbac4dd]",            // wrong braces
		"urn:uuid:",                                         // urn empty
		"urn:uuid:e8c8cec324e9445aa0",                       // urn too short
		"urn:uuid:e8c8cec324e9445aa086f021ecbac4ddaaaaaaaa", // urn too long
		"urn:uuid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",         // urn not hex
		"urn:uuid:e8c8cec3-24e944-5a-a086f021ecbac-4dd",     // urn dashes misplaced
	}
	for _, str := range testStrings {
		t.Run(str, func(t *testing.T) {
			var u UUID
			if err := u.UnmarshalText([]byte(str)); err == nil {
				t.Fatalf("expected fail parsing: %v", u)
			}
		})
	}
}

func TestUnmarshalUUID(t *testing.T) {
	expected := UUID{0xe8, 0xc8, 0xce, 0xc3, 0x24, 0xe9, 0x44, 0x5a, 0xa0, 0x86, 0xf0, 0x21, 0xec, 0xba, 0xc4, 0xdd}
	testStrings := []string{
		"e8c8cec324e9445aa086f021ecbac4dd",
		"e8c8cec3-24e9-445a-a086-f021ecbac4dd",
		"{e8c8cec3-24e9-445a-a086-f021ecbac4dd}",
		"urn:uuid:e8c8cec324e9445aa086f021ecbac4dd",
		"urn:uuid:e8c8cec3-24e9-445a-a086-f021ecbac4dd",
	}
	for _, str := range testStrings {
		t.Run(str, func(t *testing.T) {
			var u UUID
			if err := u.UnmarshalText([]byte(str)); err != nil {
				t.Fatalf("parsing: %v", err)
			}
			if u != expected {
				t.Fatalf("expected uuid '%s' got '%s'", expected, u)
			}
		})
	}
}

func TestGenerateV1(t *testing.T) {
	uuids := make(map[UUID]bool)
	for i := 0; i < 1000; i++ {
		u0, err := GenerateV1()
		if err != nil {
			t.Fatal(err)
		}
		u1, err := GenerateV1()
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := uuids[u0]; ok {
			t.Fatalf("u0: Conflict: %s", u0)
		}
		if _, ok := uuids[u1]; ok {
			t.Fatalf("u1: Conflict: %s", u0)
		}
		uuids[u0] = true
		uuids[u1] = true
		u2, err := ParseUUIDString(u1.String())
		if err != nil {
			t.Fatal(err)
		}
		if u1 != u2 {
			t.Fatalf("%s != parsed %s", u1, u2)
		}
		if u1.Version() != 1 {
			t.Errorf("invalid version '%d'. expected 1", u1.Version())
		}
		if v := u1.Variant(); v != VariantRFC4122 {
			t.Errorf("incorrect variant '%d'. expected '%d'", v, VariantRFC4122)
		}
		//t0 := u0.Time()
		//t1 := u1.Time()
		//if t1.Before(t0) {
		//	t.Errorf("time went backwards '%s' < '%s", t1, t0)
		//}
	}
}

func TestGenerateV2(t *testing.T) {
	var domain byte = 1
	var id uint32 = 5000
	uuids := make(map[UUID]bool)
	for i := 0; i < 64; i++ {
		u0, err := GenerateV2(domain, id)
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := uuids[u0]; ok {
			t.Fatalf("Conflict: %s", u0)
		}
		uuids[u0] = true
		u1, err := ParseUUIDString(u0.String())
		if err != nil {
			t.Fatal(err)
		}
		if u0 != u1 {
			t.Fatalf("%s != parsed %s", u0, u1)
		}
		if u1.Version() != 2 {
			t.Errorf("invalid version '%d'. expected 2", u1.Version())
		}
		if v := u1.Variant(); v != VariantRFC4122 {
			t.Errorf("incorrect variant '%d'. expected '%d'", v, VariantRFC4122)
		}
		d0, id0 := u0.DCESecurity()
		//t0 := u0.Time()
		//t.Logf("u0: %s - %v (%x/%d)", u0, t0, d0, id0)
		if d0 != domain {
			t.Errorf("incorrect domain '%x'. expected '%x'", d0, domain)
		}
		if id0 != id {
			t.Errorf("incorrect id '%x'. expected '%x'", id0, id)
		}
	}
}

func TestGenerateV4(t *testing.T) {
	for i := 0; i < 1000; i++ {
		u, err := GenerateV4()
		if err != nil {
			t.Fatal(err)
		}
		u2, err := ParseUUIDString(u.String())
		if err != nil {
			t.Fatal(err)
		}
		if u != u2 {
			t.Fatalf("%s != parsed %s", u, u2)
		}
		if u.Version() != 4 {
			t.Errorf("invalid version '%d'. expected 4", u.Version())
		}
		if v := u.Variant(); v != VariantRFC4122 {
			t.Errorf("incorrect variant '%d'. expected '%d'", v, VariantRFC4122)
		}
		//t.Log(u)
	}
}

func TestFormatUUID(t *testing.T) {
	u := MustParseUUIDString("9073926b-929f-31c2-abc9-fad77ae3e8eb")
	tests := []struct {
		Format   string
		Expected string
	}{
		{"%s", "9073926b-929f-31c2-abc9-fad77ae3e8eb"},
		{"%+s", "9073926B-929F-31C2-ABC9-FAD77AE3E8EB"},
		{"%v", "9073926b-929f-31c2-abc9-fad77ae3e8eb"},
		{"%+v", "9073926B-929F-31C2-ABC9-FAD77AE3E8EB"},
		{"%x", "9073926b929f31c2abc9fad77ae3e8eb"},
		{"%X", "9073926B929F31C2ABC9FAD77AE3E8EB"},
		{"%q", `"9073926b-929f-31c2-abc9-fad77ae3e8eb"`},
		{"%+q", `"9073926B-929F-31C2-ABC9-FAD77AE3E8EB"`},
	}
	for _, test := range tests {
		t.Run(test.Format, func(t *testing.T) {
			f := fmt.Sprintf(test.Format, u)
			if f != test.Expected {
				t.Errorf("got '%s', expected '%s", f, test.Expected)
			}
			//t.Logf(f)
		})
	}

}

func TestHashBasedUUID(t *testing.T) {
	tests := []struct {
		namespace UUID
		name      string
		expected3 UUID
		expected5 UUID
	}{
		{NamespaceDNS, "example.com", MustParseUUIDString("9073926b-929f-31c2-abc9-fad77ae3e8eb"), MustParseUUIDString("cfbff0d1-9375-5685-968c-48ce8b15ae17")},
		{NamespaceX500, "example.com", MustParseUUIDString("11c2f001-e3a4-3ad0-90f7-88ac418c36b8"), MustParseUUIDString("f014ed3c-177a-541e-a840-fc330670f8e8")},
		{NamespaceOID, "example.com", MustParseUUIDString("109f8204-164d-33ef-871d-d92c373e8c66"), MustParseUUIDString("eb6106fd-8a37-5395-b3f7-7cb93195fdba")},
		{NamespaceDNS, "www.example.com", MustParseUUIDString("5df41881-3aed-3515-88a7-2f4a814cf09e"), MustParseUUIDString("2ed6657d-e927-568b-95e1-2665a8aea6a2")},
		{NamespaceURL, "https://www.example.com/uuid5", MustParseUUIDString("73a6ec42-6919-32f6-95e5-ae233b1dbfb9"), MustParseUUIDString("268f0b2f-1cb0-5e48-a699-a61590854f48")},
		{NamespaceURL, "https://www.example.com/uuid5?help", MustParseUUIDString("d74cf8b7-d8ca-360c-896d-e3b1a295d1df"), MustParseUUIDString("37c80565-384a-5dde-aead-8a190c9cbf8e")},
	}
	for _, ex := range tests {
		t.Run(fmt.Sprintf("v3-%s-%s", ex.namespace, ex.name), func(t *testing.T) {
			u := GenerateV3(ex.namespace, []byte(ex.name))
			if u != ex.expected3 {
				t.Errorf("%s != %s", u, ex.expected3)
			}
			if u.Version() != 3 {
				t.Errorf("invalid version '%d'. expected 3", u.Version())
			}
			if v := u.Variant(); v != VariantRFC4122 {
				t.Errorf("incorrect variant '%d'. expected '%d'", v, VariantRFC4122)
			}
		})
		t.Run(fmt.Sprintf("v5-%s-%s", ex.namespace, ex.name), func(t *testing.T) {
			u := GenerateV5(ex.namespace, []byte(ex.name))
			if u != ex.expected5 {
				t.Errorf("%s != %s", u, ex.expected5)
			}
			if u.Version() != 5 {
				t.Errorf("invalid version '%d'. expected 5", u.Version())
			}
			if v := u.Variant(); v != VariantRFC4122 {
				t.Errorf("incorrect variant '%d'. expected '%d'", v, VariantRFC4122)
			}
		})
	}
}

// ExampleParseUUIDString parses a uuid in braced form and prints the canonical form
func ExampleParseUUIDString() {
	uuid, _ := ParseUUIDString("{e3b4fa08-0365-403d-bc0c-5a3589a1401d}")
	fmt.Println(uuid)
	// Output: e3b4fa08-0365-403d-bc0c-5a3589a1401d
}

// ExampleGenerateV4 generates a V4 UUID
func ExampleGenerateV1() {
	uuid, err := GenerateV1()
	if err != nil {
		fmt.Println("Can't generate Time-UUID:", err)
		return
	}
	fmt.Println(uuid)
}

// ExampleGenerateV4 generates a V4 UUID
func ExampleGenerateV4() {
	uuid, err := GenerateV4()
	if err != nil {
		fmt.Println("Can't generate Random UUID:", err)
		return
	}
	fmt.Println(uuid)
}
