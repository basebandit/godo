package godo

const Schema = `
  	schema {
    		query: Query
 	}
 	# The query type, represents all of the entry points into our object graph
 	type Query{
     		tasks: [Task]!
     		task(id: ID!) : Task
	}
    
	scalar Time

	type Task{
			id: ID!
			name: String
			description: String
			startDate: Time!
			endDate: Time!
			done: Boolean!
	}
`
