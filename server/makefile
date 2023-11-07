KEY_DIR:=./mock
PRIVATE_KEY_PATH:=./mock/secure/private/private.pem

.Phony: gen_key
gen_key:
	@echo "PRIVATE_KEY_PATH=$(PRIVATE_KEY_PATH)"
	@echo "KEY_DIR=$(KEY_DIR)"
	@mkdir -p $(KEY_DIR)
	@openssl genrsa -out $(PRIVATE_KEY_PATH) 2048
	@chmod 700 $(KEY_DIR)
	@chmod 600 $(PRIVATE_KEY_PATH)
	@echo "done generate key!!"
