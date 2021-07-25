cfssl print-defaults config > config.json
cfssl print-defaults csr > config.json

cfssl gencert -initca ca-csr.json | cfssljson -bare ca

cfssl gencert -ca ca/ca.pem -ca-key ca/ca-key.pem -config config.json -profile www server-csr.json|cfssljson -bare server/server