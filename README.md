```bash
# 创建项目（默认启用 Go Modules）
bee new myproject
cd myproject

# 初始化依赖
go mod tidy

bee run -gendoc=true -downdoc=true
导出为 JSON/YAML

# 生成 swagger.json
swagger generate spec -o ./swagger/swagger.json

```
