.PHONY: run-admin-svc
run-admin-svc:
	cd apps/admin-svc && air

.PHONY: run-setting-svc
run-setting-svc:
	cd apps/setting-svc && air

.PHONY: run-auth-svc
run-auth-svc:
	cd apps/auth-svc && air