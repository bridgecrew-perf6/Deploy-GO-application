package dao

var (
	ErrorFailedToUnmarshalRecord      = "failed to unmarshal record"
	ErrorFailedToFetchRecord          = "failed to fetch record"
	ErrorInvalidUserData              = "invalid user data"
	ErrorCouldNotMarshalItem          = "could not marshal item"
	ErrorCouldNotDeleteItem           = "could not delete item"
	ErrorCouldNotDynamoPutItem        = "could not dynamo put item error"
	ErrorUserAlreadyExists            = "User already exists"
	ErrorUserDoesNotExists            = "User does not exist"
	ErrorMethodNotAllowed             = "method Not allowed"
	ErrorInternalServer               = "Internal Server Error"
	ErrorCouldNotConvertStringIntoInt = "Could not convert string into int"
	ErrorQueryStringParameter         = "Fill the QueryString Parameter Correctly"
	STUDENTNOTALLOWED                 = "This tracker is available only for scoreplus paid students"
)
