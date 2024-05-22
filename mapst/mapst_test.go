package mapst_test

import (
	"testing"

	"github.com/aura-studio/boost/mapst"
	. "github.com/frankban/quicktest"
)

func TestConvertMapStringInterfaceToStruct(t *testing.T) {
	c := New(t)

	c.Run("ConvertMapStringInterfaceToStruct", func(c *C) {
		src := map[string]interface{}{
			"Name": "John",
			"Age":  30,
		}

		dst := struct {
			Name string
			Age  int
		}{}

		mapst.Convert(src, &dst)

		c.Assert(dst.Name, Equals, "John")
		c.Assert(dst.Age, Equals, 30)
	})
}

func TestConvertStructToMapStringInterface(t *testing.T) {
	c := New(t)

	c.Run("ConvertStructToMapStringInterface", func(c *C) {
		src := struct {
			Name string
			Age  int
		}{
			Name: "John",
			Age:  30,
		}

		dst := map[string]interface{}{}

		mapst.Convert(&src, &dst)

		c.Assert(dst["Name"], Equals, "John")
		c.Assert(dst["Age"], Equals, 30)
	})
}

func TestConvertUnsupportedType(t *testing.T) {
	c := New(t)

	c.Run("ConvertUnsupportedType", func(c *C) {
		src := 1
		dst := 2

		err := mapst.ConvertE(src, &dst)

		c.Assert(err, ErrorMatches, "type conversion not supported")
	})
}

func TestConvertMapStringInterfaceToStructDeep(t *testing.T) {
	c := New(t)

	c.Run("ConvertMapStringInterfaceToStructDeep", func(c *C) {
		src := map[string]interface{}{
			"Name": "John",
			"Age":  30,
			"Address": map[string]interface{}{
				"City": "Jakarta",
				"Zip":  12345,
			},
		}

		dst := struct {
			Name    string
			Age     int
			Address struct {
				City string
				Zip  int
			}
		}{}

		mapst.Convert(src, &dst)

		c.Assert(dst.Name, Equals, "John")
		c.Assert(dst.Age, Equals, 30)
		c.Assert(dst.Address.City, Equals, "Jakarta")
		c.Assert(dst.Address.Zip, Equals, 12345)
	})
}

func TestConvertStructToMapStringInterfaceDeep(t *testing.T) {
	c := New(t)
	c.Run("ConvertStructToMapStringInterfaceDeep", func(c *C) {
		src := struct {
			Name    string
			Age     int
			Address struct {
				City string
				Zip  int
			}
		}{
			Name: "John",
			Age:  30,
			Address: struct {
				City string
				Zip  int
			}{
				City: "Jakarta",
				Zip:  12345,
			},
		}

		dst := map[string]interface{}{}

		mapst.Convert(&src, &dst)

		c.Assert(dst["Name"], Equals, "John")
		c.Assert(dst["Age"], Equals, 30)
		c.Assert(dst["Address"].(map[string]interface{})["City"], Equals, "Jakarta")
		c.Assert(dst["Address"].(map[string]interface{})["Zip"], Equals, 12345)
	})
}

func TestConvertMapStringInterfaceToStructDeepWithSlice(t *testing.T) {
	c := New(t)

	c.Run("ConvertMapStringInterfaceToStructDeepWithSlice", func(c *C) {
		src := map[string]interface{}{
			"Name": "John",
			"Age":  30,
			"Addresses": []map[string]interface{}{
				{
					"City": "Jakarta",
					"Zip":  12345,
				},
				{
					"City": "Bandung",
					"Zip":  67890,
				},
			},
		}

		dst := struct {
			Name      string
			Age       int
			Addresses []struct {
				City string
				Zip  int
			}
		}{}

		mapst.Convert(src, &dst)

		c.Assert(dst.Name, Equals, "John")
		c.Assert(dst.Age, Equals, 30)
		c.Assert(dst.Addresses[0].City, Equals, "Jakarta")
		c.Assert(dst.Addresses[0].Zip, Equals, 12345)
		c.Assert(dst.Addresses[1].City, Equals, "Bandung")
		c.Assert(dst.Addresses[1].Zip, Equals, 67890)
	})
}

func TestConvertStructToMapStringInterfaceDeepWithSlice(t *testing.T) {
	c := New(t)

	c.Run("ConvertStructToMapStringInterfaceDeepWithSlice", func(c *C) {
		src := struct {
			Name      string
			Age       int
			Addresses []struct {
				City string
				Zip  int
			}
		}{
			Name: "John",
			Age:  30,
			Addresses: []struct {
				City string
				Zip  int
			}{
				{
					City: "Jakarta",
					Zip:  12345,
				},
				{
					City: "Bandung",
					Zip:  67890,
				},
			},
		}

		dst := map[string]interface{}{}

		mapst.Convert(&src, &dst)

		c.Assert(dst["Name"], Equals, "John")
		c.Assert(dst["Age"], Equals, 30)
		c.Assert(dst["Addresses"].([]interface{})[0].(map[string]interface{})["City"], Equals, "Jakarta")
		c.Assert(dst["Addresses"].([]interface{})[0].(map[string]interface{})["Zip"], Equals, 12345)
		c.Assert(dst["Addresses"].([]interface{})[1].(map[string]interface{})["City"], Equals, "Bandung")
		c.Assert(dst["Addresses"].([]interface{})[1].(map[string]interface{})["Zip"], Equals, 67890)
	})
}
