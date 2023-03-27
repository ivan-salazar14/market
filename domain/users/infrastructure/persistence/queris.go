package persistence

const (
	InsertUser  = "INSERT INTO public.users (user_id,user_name,user_amount,user_user_created,user_date_created,user_user_modify,user_date_modify) VALUES ($1,$2,$3,$4,'NOW()',$5,'NOW()')"
	SelectUser  = "SELECT user_id, user_name, user_amount, user_user_created, user_date_created, user_user_modify, user_date_modify FROM public.users WHERE user_id = $1"
	SelectUsers = "SELECT user_id, user_name, user_amount, user_user_created, user_date_created, user_user_modify, user_date_modify FROM public.users"
)
