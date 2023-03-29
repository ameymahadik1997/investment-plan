json.extract! user_detail, :id, :first_name, :last_name, :age, :career, :input_income, :created_at, :updated_at
json.url user_detail_url(user_detail, format: :json)
