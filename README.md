# <app-name>

## This project is generated from template [t-ddd-fiber-gorm](https://github.com/loopopenen/t-ddd-fiber-gorm).

## 1. Use make.
* Use make to rename application name.
```sh
make name org=<org-name> app=<app-name>
```

* Use make to build swagger documents.
```sh
make doc
```

* Use make to run your project.
```sh
make run

make run env=local
```

* Use make to wire dependency injection
```sh
make wire
```

* Use make to generate all.
```sh
make gen
```

## 2. Run and access http://127.0.0.1:8080/swagger/index.html

## 3. Delete file PLEASE_DELETE_ME.go and fix your code.

## 4. Create a new git repository of the command line:
```sh
git init
git add -A
git commit -m "init commit"
git branch -M main
git remote add origin git@github.com:<org-name>/<app-name>.git
git push -u origin main
```

---

* ## How to use gonew template?
### 1. Install gonew.
```sh
go install golang.org/x/tools/cmd/gonew@latest
```
### 2. Use this template to create your project.
```sh
mkdir <app-name> && \
gonew github.com/loopopen/t-ddd-fiber-gorm github.com/<org-name>/<app-name> ./<app-name>
```

