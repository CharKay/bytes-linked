package e

import (
	"net/http"
	"strconv"
)

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400

	ERROR_FILE_GET_FORM_FAIL = 40001
	ERROR_USER_PASSWORD      = 40002
	ERROR_USER_TOKEN         = 40003

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004

	ERROR_FILE_GET_FILEMETA_NIL                 = 22001
	ERROR_FILE_SAVE_FILEMETA_FAIL               = 22002
	ERROR_FILE_GET_FILEMETA_FAIL                = 22003
	ERROR_FILE_UPLOADED_SAVE_FILEMETA_FAIL      = 22004
	ERROR_FILE_QUERY_USERFILEMETA_FAIL          = 22005
	ERROR_FILE_SAVE_USERFILEMETA_FAIL           = 22006
	ERROR_FILE_DELETE_USERFILE_FAIL             = 22007
	ERROR_FILE_DELETE_USERFILE_FAIL_FOR_NOOWNER = 22008

	ERROR_USER_SIGNUP_FAIL              = 50001
	ERROR_USER_GET_INFO_FAIL            = 50002
	ERROR_USER_CHECK_EXIST_FAIL         = 50003
	ERROR_USER_ALREADY_EXIST            = 50004
	ERROR_USER_IS_NOT_ADMIN             = 50005
	ERROR_USER_QUERY_FAIL               = 50006
	ERROR_USER_UPDATE_PASSWORD_FAIL     = 50007
	ERROR_USER_DELETE_FAIL              = 50008
	ERROR_USER_ADMIN_TOKEN              = 50009
	ERROR_USER_INVITE_CODE              = 50010
	ERROR_USER_SECRET_FILE_QUERY_FAILED = 50011
	ERROR_USER_PARAMS_ERROR             = 50012
	ERROR_USER_PHONE_INVALID            = 50013
	ERROR_USER_TRANSACTION_ERROR        = 50014
	ERROR_USER_INVALID_ITEM             = 50015
	ERROR_USER_UPDATE_AVATAR_FAIL       = 50016
	ERROR_USER_FILE_NOT_FOUND            = 50017

	ERROR_FILE_COMMIT_FAIL         = 51000
	ERROR_FILE_UPLOAD_TO_CEPH_FAIL = 51001
	ERROR_FILE_UPLOAD_LOCAL_FAIL   = 51002
	ERROR_FILE_UPLOAD_FAIL         = 51003
	ERROR_FILE_DOWNLOAD_ROOTDIR    = 51004
	ERROR_FILE_DELETE_FAIL         = 51005
	ERROR_FILE_RENAME_FAIL         = 51006
	ERROR_FILE_CHANGE_TO_BYTE_FAIL = 51007
	ERROR_FILE_DELETE_FILE_FAIL    = 51008

	ERROR_FILE_SHARE_FAIL               = 53002
	ERROR_FILE_CANCEL_SHARE_FAIL        = 53003
	ERROR_FILE_AUDIT_SHARE_FAIL         = 53004
	ERROR_FILE_CANCEL_AUDIT_SHARE_FAIL  = 53005
	ERROR_FILE_UNSHARE_DOWNLOAD_FAIL    = 53006
	ERROR_FILE_UNPERMISSION_DELETE_FILE = 53007

	ERROR_FILE_CANCEL_MPUPLOAD_FAIL     = 54001
	ERROR_FILE_QUERY_MPUPLOAD_FAIL      = 54002
	ERROR_FILE_UPLOAD_MPUPLOAD_FAIL     = 54003
	ERROR_FILE_COMPLETE_MPUPLOAD_FAIL   = 54004
	ERROR_FILE_INVALID_MPUPLOAD_REQUEST = 54005

	ERROR_FILE_NORECORD_FASTUPLOAD_FAIL = 55001
	ERROR_FILE_UPLOAD_FASTUPLOAD_FAIL   = 55002
)

func GetHttpCode(code int) int {
	if code == 0 {
		return http.StatusOK
	}

	codeStr := strconv.Itoa(code)
	if len(codeStr) <= 1 {
		return http.StatusOK
	}
	switch codeStr[0:1] {
	case "3":
		return http.StatusMultipleChoices
	case "4":
		return http.StatusBadRequest
	case "5":
		return http.StatusInternalServerError
	default:
		return http.StatusOK
	}
}
