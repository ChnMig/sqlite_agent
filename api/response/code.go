package response

// https://google-cloud.gitbook.io/api-design-guide/errors

type responseData struct {
	Code      int         `json:"code"`
	Status    string      `json:"status"`
	Message   string      `json:"message,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Timestamp int64       `json:"timestamp"`
	Count     *int        `json:"count,omitempty"`
}

// Common error messages

// No error.
var Success = responseData{
	Code:    200,
	Status:  "OK",
	Message: "Request successful",
}

// The data sent by the client contains invalid parameters. Check the error message and details for more information.
var INVALID_ARGUMENT = responseData{
	Code:    400,
	Status:  "INVALID_ARGUMENT",
	Message: "Invalid request parameters",
}

// The current system state does not allow the execution of the request, such as deleting a non-empty directory.
var FAILED_PRECONDITION = responseData{
	Code:    400,
	Status:  "FAILED_PRECONDITION",
	Message: "Unable to execute client request",
}

// The client specified an invalid range.
var OUT_OF_RANGE = responseData{
	Code:    400,
	Status:  "OUT_OF_RANGE",
	Message: "Client out-of-bounds access",
}

// The request failed authentication due to a missing, invalid, or expired OAuth token.
var UNAUTHENTICATED = responseData{
	Code:    401,
	Status:  "UNAUTHENTICATED",
	Message: "Authentication failed",
}

// The client does not have sufficient permissions. This could be due to an OAuth token with incorrect scope, lack of client permissions, or the API being disabled for the client code.
var PERMISSION_DENIED = responseData{
	Code:    403,
	Status:  "PERMISSION_DENIED",
	Message: "Insufficient client permissions",
}

// The specified resource was not found or the request was denied for undisclosed reasons (e.g., whitelist).
var NOT_FOUND = responseData{
	Code:    404,
	Status:  "NOT_FOUND",
	Message: "Resource not found",
}

// Concurrent conflict, such as read-modify-write conflict.
var ABORTED = responseData{
	Code:    409,
	Status:  "ABORTED",
	Message: "Data processing conflict",
}

// The resource the client attempted to create already exists.
var ALREADY_EXISTS = responseData{
	Code:    409,
	Status:  "ALREADY_EXISTS",
	Message: "Resource already exists",
}

// Resource quota exceeded or rate limit reached.
var RESOURCE_EXHAUSTED = responseData{
	Code:    429,
	Status:  "RESOURCE_EXHAUSTED",
	Message: "Resource quota exceeded or rate limit reached",
}

// The request was cancelled by the client.
var CANCELLED = responseData{
	Code:    499,
	Status:  "CANCELLED",
	Message: "Request cancelled by client",
}

// Irrecoverable data loss or corruption. The client should report the error to the user.
var DATA_LOSS = responseData{
	Code:    500,
	Status:  "DATA_LOSS",
	Message: "Error processing data",
}

// Unknown server error, usually due to a server bug.
var UNKNOWN = responseData{
	Code:    500,
	Status:  "UNKNOWN",
	Message: "Unknown server error",
}

// Internal server error, usually due to a server bug.
var INTERNAL = responseData{
	Code:    500,
	Status:  "INTERNAL",
	Message: "Internal server error",
}

// The API method is not implemented by the server.
var NOT_IMPLEMENTED = responseData{
	Code:    501,
	Status:  "NOT_IMPLEMENTED",
	Message: "API not implemented",
}

// Service unavailable, usually due to server downtime.
var UNAVAILABLE = responseData{
	Code:    503,
	Status:  "UNAVAILABLE",
	Message: "Service unavailable",
}

// The request exceeded the deadline. This occurs only when the caller sets a deadline shorter than the method's default deadline and the request does not complete within the deadline.
var DEALINE_EXCEED = responseData{
	Code:    504,
	Status:  "DEALINE_EXCEED",
	Message: "Request timeout",
}

// Custom status codes defined by business logic
