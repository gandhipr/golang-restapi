package messages

var (
	RequiredError = "The field is required, cannot be empty"
	EmailError    = "The field should be a valid email"
	TitleError    = "The field should be a valid title, allowed chars: [alpha, num, [.]dot, [-]hyphen, [_]underscore, [/]backslash, [,]comma, [ ]space]"
	VersionError  = "The field should be a valid version, allowed chars: [alpha, num, [.]dot, [-]hyphen, [_]underscore, [/]backslash]"
	DefaultError  = "The field is invalid"
)
