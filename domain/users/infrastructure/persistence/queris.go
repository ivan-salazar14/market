package persistence

const (
	InsertUser  = "INSERT INTO public.users (user_id,user_name,user_identifier,user_email,user_password,user_type_identifier,user_date_modify) VALUES ($1,$2,$3,$4,$5,$6,'NOW()')"
	SelectUser  = "SELECT user_id, user_name, user_identifier, user_email, user_password, user_type_identifier, user_date_modify FROM public.users WHERE user_id = $1"
	SelectUsers = "SELECT user_id, user_name, user_identifier, user_email, user_password, user_type_identifier, user_date_modify FROM public.users"
)
