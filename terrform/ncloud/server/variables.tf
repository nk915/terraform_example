variable "access_key" {
	sensitive = true
	description = "사용자 접근 키"
}
variable "secret_key" {
	sensitive = true
	description = "사용자 비밀 키"
}
variable "region" {
	default = "KR"
}
