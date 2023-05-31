package errno

const (
	Success = 200

	UnKnowError   = 99999
	ServerError   = 500
	ParamsInvalid = 40001
	CallRPCFailed = 40002

	UserCreateFail    = 10101
	UserLoginFail     = 10102
	CheckPasswordFail = 10103
	UserNoexist       = 10104

	TaskCreateFail  = 10201
	TaskUpdateFail  = 10202
	GetTaskListFail = 10203
	TaskDeleteFail  = 10204
	TaskSearchFail  = 10205

	TokenGenerateFail = 10301
	TokenParseFail    = 10302
	TokenExpired      = 10303
	NoToken           = 10304
)

var CodeTag = map[int]string{
	Success:     "success",
	ServerError: "server internal error",

	ParamsInvalid: "Params Invalid",
	CallRPCFailed: "rpc called failed",

	UnKnowError:       "UnKnow error",
	UserCreateFail:    "User Create Fail",
	UserLoginFail:     "User Login Fail",
	CheckPasswordFail: "Username or Password error",
	UserNoexist:       "User No Exist",

	TaskCreateFail:  "Task Create Fail",
	TaskUpdateFail:  "Task Update Fail",
	GetTaskListFail: "Get TaskList Fail",
	TaskDeleteFail:  "Task Delete Fail",
	TaskSearchFail:  "Task Search Fail",

	TokenGenerateFail: "Token Generate Fail",
	TokenParseFail:    "Token Parse Fail",
	TokenExpired:      "Token Expired",
	NoToken:           "No Token",
}
