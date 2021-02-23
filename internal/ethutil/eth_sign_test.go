package ethutil

import "testing"

func TestVerify(t *testing.T) {
	data := []byte(`{"abstract":"","category":"","content_hash":"2138e51879953bcb3870b914f888a97dce283414ca7ad4283b83e47153584b4b","created":1543585869000,"creator":{"account":{"address":"","avatar":{"abstract":"","category":"","content_hash":"06c13a8acf379ab6a12246e40e4889a9fdcb49215065e30252e5f382142e1dd3","created":1545137474000,"extra":{"dtcp_id":"59219e7912ddb34eab2baa8f348c5b1fcb199f290df0aea57783d7a462047fe0"},"language":"","status":"created","tag":"image","title":"","type":"object","version":"1.0"},"created":1546416086277,"creator":{"account_id":"7139407a2e15c171ebb971c0aeebbcbf33197c448b34d92072248b25d435398c","sub_account_id":""},"extra":{"hash":""},"name":"富贵闲人","status":"created","tag":"account","type":"object","version":"1.0"}},"extra":{"dtcp_id":"d52e8050341350c483d530a44032f29711a66ca29e7d2834c197dc09ccce4564","group_dtcp_id":"76484c165999720eb0c614c9ed317e8e5aa68bb558b7cfdcd34c4af969f1f08c","images":null,"link":null},"language":"","status":"created","tag":"article","title":"","type":"object","version":"1.0"}`)
	signature := "1cc2a25209bc28732d0c4e7c458aac2839598f91a14ee93f9b79e0527998fcbe7da5f1e961e9c82cdf81a455112ca4205ace8932ee291c10b6f27efe176a5f5000"
	nodeAddress := "0x5bb297a46512233e9f52f74e0cafd6ecb2d2db07"

	if !Verify(data, signature, nodeAddress) {
		t.Errorf("check error")
	}
}
