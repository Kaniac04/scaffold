package templates

var AvailableProjects = map[string]struct{}{
	"flask":  {},
	"django": {},
	"fast":   {},
	"gin":    {},
	"fiber":  {},
	"echo":   {},
}

var ProjectDirectories = map[string][]string{
	"flask": {
		"flask/app",
		"flask/app/templates",
		"flask/app/static",
		"flask/app/__init__.py",
		"flask/app/routes.py",
		"flask/app/models.py",
		"flask/app/services.py",
		"flask/app/utils.py",
		"flask/migrations",
		"flask/tests",
		"flask/requirements.txt",
		"flask/run.py",
	},
	"django": {
		"django/my_django_project",
		"django/my_django_project/settings.py",
		"django/my_django_project/urls.py",
		"django/my_django_project/wsgi.py",
		"django/my_django_project/asgi.py",
		"django/apps/main_app",
		"django/apps/main_app/migrations",
		"django/apps/main_app/__init__.py",
		"django/apps/main_app/models.py",
		"django/apps/main_app/views.py",
		"django/apps/main_app/urls.py",
		"django/apps/main_app/forms.py",
		"django/static",
		"django/media",
		"django/tests",
		"django/manage.py",
		"django/requirements.txt",
	},
	"fast": {
		"fast/app",
		"fast/app/__init__.py",
		"fast/app/routes.py",
		"fast/app/models.py",
		"fast/app/services.py",
		"fast/app/utils.py",
		"fast/app/templates",
		"fast/app/static",
		"fast/migrations",
		"fast/tests",
		"fast/requirements.txt",
		"fast/run.py",
	},
	"gin": {
		"gin-app/cmd/main.go",
		"gin-app/config/config.go",
		"gin-app/internal/controllers/user_controller.go",
		"gin-app/internal/controllers/auth_controller.go",
		"gin-app/internal/models/user.go",
		"gin-app/internal/routes/routes.go",
		"gin-app/internal/middlewares/auth_middleware.go",
		"gin-app/internal/services/auth_service.go",
		"gin-app/internal/utils/logger.go",
		"gin-app/internal/utils/response.go",
		"gin-app/database/db.go",
		"gin-app/database/migrations/",
		"gin-app/tests/controllers/",
		"gin-app/tests/routes/",
		"gin-app/.env",
		"gin-app/.gitignore",
		"gin-app/go.mod",
		"gin-app/go.sum",
		"gin-app/README.md",
	},
	"fiber": {
		"fiber-app/cmd/main.go",
		"fiber-app/config/config.go",
		"fiber-app/internal/controllers/user_controller.go",
		"fiber-app/internal/controllers/auth_controller.go",
		"fiber-app/internal/models/user.go",
		"fiber-app/internal/routes/routes.go",
		"fiber-app/internal/middlewares/auth_middleware.go",
		"fiber-app/internal/services/auth_service.go",
		"fiber-app/internal/utils/logger.go",
		"fiber-app/internal/utils/response.go",
		"fiber-app/database/db.go",
		"fiber-app/database/migrations/",
		"fiber-app/tests/controllers/",
		"fiber-app/tests/routes/",
		"fiber-app/.env",
		"fiber-app/.gitignore",
		"fiber-app/go.mod",
		"fiber-app/go.sum",
		"fiber-app/README.md",
	},
	"echo": {
		"echo-app/cmd/main.go",
		"echo-app/config/config.go",
		"echo-app/internal/controllers/user_controller.go",
		"echo-app/internal/controllers/auth_controller.go",
		"echo-app/internal/models/user.go",
		"echo-app/internal/routes/routes.go",
		"echo-app/internal/middlewares/auth_middleware.go",
		"echo-app/internal/services/auth_service.go",
		"echo-app/internal/utils/logger.go",
		"echo-app/internal/utils/response.go",
		"echo-app/database/db.go",
		"echo-app/database/migrations/",
		"echo-app/tests/controllers/",
		"echo-app/tests/routes/",
		"echo-app/.env",
		"echo-app/.gitignore",
		"echo-app/go.mod",
		"echo-app/go.sum",
		"echo-app/README.md",
	},
}

var ProjectTemplates = map[string]string{
	"flask": `
flask/
├── app/
│   ├── templates/
│   ├── static/
│   ├── __init__.py
│   ├── routes.py
│   ├── models.py
│   ├── services.py
│   ├── utils.py
├── migrations/
├── tests/
├── requirements.txt
└── run.py
`,
	"django": `
django/
├── my_django_project/
│   ├── settings.py
│   ├── urls.py
│   ├── wsgi.py
│   ├── asgi.py
├── apps/
│   └── main_app/
│       ├── __init__.py
│       ├── models.py
│       ├── views.py
│       ├── urls.py
│       ├── forms.py
│       └── migrations/
├── static/
├── media/
├── manage.py
├── requirements.txt
└── .env
`,
	"fast": `
fast/
├── app/
│   ├── templates/
│   ├── static/
│   ├── __init__.py
│   ├── routes.py
│   ├── models.py
│   ├── services.py
│   ├── utils.py
├── migrations/
├── tests/
├── requirements.txt
└── run.py
`,
	"gin": `
gin-app/
│── cmd/
│   ├── main.go
│
│── config/
│   ├── config.go
│
│── internal/
│   ├── controllers/
│   │   ├── user_controller.go
│   │   ├── auth_controller.go
│   │
│   ├── models/
│   │   ├── user.go
│   │
│   ├── routes/
│   │   ├── routes.go
│
│   ├── middlewares/
│   │   ├── auth_middleware.go
│
│   ├── services/
│   │   ├── auth_service.go
│
│   ├── utils/
│   │   ├── logger.go
│   │   ├── response.go
│
│── database/
│   ├── db.go
│   ├── migrations/
│
│── tests/
│   ├── controllers/
│   ├── routes/
│
│── .env
│── .gitignore
│── go.mod
│── go.sum
│── README.md
`,
	"fiber": `
fiber-app/
│── cmd/
│   ├── main.go
│
│── config/
│   ├── config.go
│
│── internal/
│   ├── controllers/
│   │   ├── user_controller.go
│   │   ├── auth_controller.go
│   │
│   ├── models/
│   │   ├── user.go
│   │
│   ├── routes/
│   │   ├── routes.go
│
│   ├── middlewares/
│   │   ├── auth_middleware.go
│
│   ├── services/
│   │   ├── auth_service.go
│
│   ├── utils/
│   │   ├── logger.go
│   │   ├── response.go
│
│── database/
│   ├── db.go
│   ├── migrations/
│
│── tests/
│   ├── controllers/
│   ├── routes/
│
│── .env
│── .gitignore
│── go.mod
│── go.sum
│── README.md
`,
	"echo": `
echo-app/
│── cmd/
│   ├── main.go
│
│── config/
│   ├── config.go
│
│── internal/
│   ├── controllers/
│   │   ├── user_controller.go
│   │   ├── auth_controller.go
│   │
│   ├── models/
│   │   ├── user.go
│
│   ├── routes/
│   │   ├── routes.go
│
│   ├── middlewares/
│   │   ├── auth_middleware.go
│
│   ├── services/
│   │   ├── auth_service.go
│
│   ├── utils/
│   │   ├── logger.go
│   │   ├── response.go
│
│── database/
│   ├── db.go
│   ├── migrations/
│
│── tests/
│   ├── controllers/
│   ├── routes/
│
│── .env
│── .gitignore
│── go.mod
│── go.sum
│── README.md
`,
}
