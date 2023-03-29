class CreateUserDetails < ActiveRecord::Migration[7.0]
  def change
    create_table :user_details do |t|
      t.string :first_name
      t.string :last_name
      t.string :age
      t.string :career
      t.integer :input_income

      t.timestamps
    end
  end
end
