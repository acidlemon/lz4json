package lz4json

import (
	"fmt"
	"testing"
)

func TestInteroperability(t *testing.T) {
	// target created by Compress::LZ4 (perl)
	target := []byte{
		0x34, 0x00, 0x00, 0x00, 0xaf, 0x7b, 0x22, 0x70,
		0x6f, 0x77, 0x61, 0x77, 0x61, 0x22, 0x3a, 0x0a,
		0x00, 0x0b, 0xc0, 0x22, 0x28, 0x27, 0x2d, 0x27,
		0x2a, 0x29, 0x22, 0x7d, 0x7d, 0x7d, 0x7d,
	}

	container := map[string]interface{}{}

	err := Unmarshal(target, &container)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if v1, ok := container["powawa"]; !ok {
		t.Errorf("container does not contains powawa")
	} else {
		c1 := v1.(map[string]interface{})
		if v2, ok := c1["powawa"]; !ok {
			t.Errorf("v1 does not contains powawa")
		} else {
			c2 := v2.(map[string]interface{})
			if v3, ok := c2["powawa"]; !ok {
				t.Errorf("v2 does not contains powawa")
			} else {
				c3 := v3.(map[string]interface{})
				if v4, ok := c3["powawa"]; !ok {
					t.Errorf("v3 does not contains powawa")
				} else {
					if v4 != "('-'*)" {
						t.Errorf(`(#'-') < v4 is not "('-'*)"`)
					} else {
						fmt.Println(container)
					}
				}
			}
		}
	}
}

func TestMarshal(t *testing.T) {
	v := map[string]interface{}{
		"text":   "ぽわわ",
		"number": float64(42),
		"object": map[string]interface{}{
			"key": "value",
		},
	}

	data, err := Marshal(v)
	if err != nil {
		t.Fatalf(err.Error())
	}

	result := map[string]interface{}{}
	err = Unmarshal(data, &result)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if result["text"] != v["text"] {
		t.Fatalf(`v.text is unmatch: result["text"]=%s, v["text"]=%s`,
			result["text"], v["text"])
	}

	if result["number"] != v["number"] {
		t.Fatalf(`v.number is unmatch: result["number"]=%s, v["number"]=%s`,
			result["number"], v["number"])
	}

	if result["object"].(map[string]interface{})["key"] !=
		v["object"].(map[string]interface{})["key"] {
		t.Fatal(`v.object.key is unmatch`)
	}

}

/*
use strict;
use warnings;
use utf8;

use 5.014;

use Compress::LZ4;
use JSON::XS;
use Encode;


sub test {
    my $hash = shift;

    my $json = JSON::XS->new->utf8->canonical(1);
    my $jsstr = $json->encode($hash);

    say $jsstr;

    my $lz4str = compress(encode_utf8($jsstr));
    my @bin = split //, $lz4str;

    my $i = 0;
    for my $c (@bin) {
        print '0x' . unpack("H2", $c) . ' ' ;
        $i++;

        if ($i % 8 == 0) {
            print "\n";
        }
    }

    print "\n";
}


test({
    text => 'ぽわわ',
    number => 42,
    object => {
        key => 'value',
    },
});

test({
    powawa => {
        powawa => {
            powawa => {
                powawa => "('-'*)",
            },
        },
    },
});

__END__
$ perl comp.pl
{"number":42,"object":{"key":"value"},"text":"ぽわわ"}
0x42 0x00 0x00 0x00 0xf0 0x33 0x7b 0x22
0x6e 0x75 0x6d 0x62 0x65 0x72 0x22 0x3a
0x34 0x32 0x2c 0x22 0x6f 0x62 0x6a 0x65
0x63 0x74 0x22 0x3a 0x7b 0x22 0x6b 0x65
0x79 0x22 0x3a 0x22 0x76 0x61 0x6c 0x75
0x65 0x22 0x7d 0x2c 0x22 0x74 0x65 0x78
0x74 0x22 0x3a 0x22 0xc3 0xa3 0xc2 0x81
0xc2 0xbd 0xc3 0xa3 0xc2 0x82 0xc2 0x8f
0xc3 0xa3 0xc2 0x82 0xc2 0x8f 0x22 0x7d

{"powawa":{"powawa":{"powawa":{"powawa":"('-'*)"}}}}
0x34 0x00 0x00 0x00 0xaf 0x7b 0x22 0x70
0x6f 0x77 0x61 0x77 0x61 0x22 0x3a 0x0a
0x00 0x0b 0xc0 0x22 0x28 0x27 0x2d 0x27
0x2a 0x29 0x22 0x7d 0x7d 0x7d 0x7d

*/
