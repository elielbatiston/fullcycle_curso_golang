https://gqlgen.com/

go run github.com/99designs/gqlgen init

http://localhost:8080/

go run github.com/99designs/gqlgen generate

======================
mutation createCategory {
	createCategory(input:{name:"Tecnologia", description:"Cursos de Tecnologia"}){
    id
    name
    description
  }
}

mutation createCourse {
	createCourse(input:{name:"Full Cycle", description:"The best!", categoryId: "793aa00c-58c8-4edd-8672-cd7533367878"}){
    id
    name
    description
  }
}

query queryCategories {
  categories {
    id
    name
    description
  }
}

query queryCourses {
  courses {
    id
    name
    description
  }
}

query queryCategoriesWithCourses {
  categories {
    id
    name
    description
    courses {
      id
      name
    }
  }
}

query queryCoursesWithCategory {
  courses {
    id
    name
    description
    category {
      id
      name
      description
    }
  }
}
======================