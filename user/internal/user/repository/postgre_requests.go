package userRepository

const AddUserPostgreRequest = "INSERT INTO users VALUES (default, $1, $2,$3,$4) RETURNING user_id"

const GetUserByNamePostgreRequest = "SELECT user_id,username,email,password,concat($2::varchar,avatar),modRule FROM users WHERE username=$1"

const GetUserByIDPostgreRequest = "SELECT user_id,username,email,password,concat($2::varchar,avatar),modRule FROM users WHERE user_id=$1"

const UpdateUserPostgreRequest = "UPDATE users SET username=$2,email=$3 WHERE user_id=$1"

const UpdateEmailPostgreRequest = "UPDATE users SET email=$2 WHERE user_id=$1"

const UpdateUserAvatarPostgreRequest = "UPDATE users SET avatar=$2 WHERE user_id=$1"

const UpdateUserPasswordPostgreRequest = "UPDATE users SET  password=$2 WHERE user_id=$1"
