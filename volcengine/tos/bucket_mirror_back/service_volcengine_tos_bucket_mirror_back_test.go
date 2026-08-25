package tos_bucket_mirror_back

import (
	"reflect"
	"testing"
)

func TestNormalizePrivateSourceResponse(t *testing.T) {
	credential := map[string]interface{}{"Role": "mirror-role"}
	primary := map[string]interface{}{
		"Endpoint":           "https://tos.example.com",
		"BucketName":         "source-bucket",
		"CredentialProvider": credential,
	}
	redirect := map[string]interface{}{
		"PrivateSource": map[string]interface{}{
			"SourceEndpoint": map[string]interface{}{
				"Primary": []interface{}{primary},
			},
		},
	}

	normalizePrivateSourceResponse(redirect)

	privateSources, ok := redirect["PrivateSource"].([]interface{})
	if !ok || len(privateSources) != 1 {
		t.Fatalf("private source was not normalized: %#v", redirect["PrivateSource"])
	}
	sourceEndpoints := privateSources[0].(map[string]interface{})["SourceEndpoint"].([]interface{})
	gotPrimary := sourceEndpoints[0].(map[string]interface{})["Primary"].([]interface{})[0].(map[string]interface{})
	credentials, ok := gotPrimary["CredentialProvider"].([]interface{})
	if !ok || len(credentials) != 1 || !reflect.DeepEqual(credentials[0], credential) {
		t.Fatalf("credential provider was not normalized: %#v", gotPrimary["CredentialProvider"])
	}
}

func TestNormalizePrivateSourceResponsePrunesEmptySiblings(t *testing.T) {
	redirect := map[string]interface{}{
		"PublicSource": map[string]interface{}{
			"FixedEndpoint": false,
			"SourceEndpoint": map[string]interface{}{
				"Primary":  []interface{}{},
				"Follower": []interface{}{},
			},
		},
		"Transform": map[string]interface{}{
			"ReplaceKeyPrefix": map[string]interface{}{},
		},
		"MirrorHeader": map[string]interface{}{
			"PassAll": false,
			"Pass":    []interface{}{},
			"Remove":  []interface{}{},
		},
	}

	normalizePrivateSourceResponse(redirect)

	if _, ok := redirect["PublicSource"]; ok {
		t.Fatalf("empty public source was retained: %#v", redirect["PublicSource"])
	}
	if _, ok := redirect["Transform"]; ok {
		t.Fatalf("empty transform was retained: %#v", redirect["Transform"])
	}
	mirrorHeader := redirect["MirrorHeader"].(map[string]interface{})
	if _, ok := mirrorHeader["Pass"]; ok {
		t.Fatalf("empty header pass list was retained: %#v", mirrorHeader)
	}
	if _, ok := mirrorHeader["Remove"]; ok {
		t.Fatalf("empty header remove list was retained: %#v", mirrorHeader)
	}
}

func TestNormalizePrivateSourceResponseKeepsConfiguredSiblings(t *testing.T) {
	redirect := map[string]interface{}{
		"PublicSource": map[string]interface{}{
			"SourceEndpoint": map[string]interface{}{
				"Primary": []interface{}{"https://packages.example.com/"},
			},
		},
		"Transform": map[string]interface{}{
			"ReplaceKeyPrefix": map[string]interface{}{
				"KeyPrefix":   "cache/",
				"ReplaceWith": "deb/",
			},
		},
	}

	normalizePrivateSourceResponse(redirect)

	if _, ok := redirect["PublicSource"]; !ok {
		t.Fatal("configured public source was removed")
	}
	if _, ok := redirect["Transform"]; !ok {
		t.Fatal("configured transform was removed")
	}
}
