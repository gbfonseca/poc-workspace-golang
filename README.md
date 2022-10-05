## POC Workspace de Miniapps em GO

Para rodar é necessário instalar o GO na máquina e executar o seguinte comando dentro do diretório file-sync

```go
  go run *.go
```

Endpoints disponíveis para teste:

    - /health
    - /setup      # Necessário seu email no arquivo ame.conf.js
    - /uploadAll  # Necessário um formData com chave valor file e zip do miniapp e email com seu email 
    - /fileupload # Necessário um formData com chave valor file e arquivo modificado do miniapp e email com seu email 
    - /storage/:email/react-proj/public/index.html ## caso queira visualizar o miniapp

