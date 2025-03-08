

## Sistema de Cálculo de Rescisão Trabalhista

<p align="center">
  <img src="https://raw.githubusercontent.com/pedropamn/SeveranceCalculator/refs/heads/main/screen.png" alt="Imagem de exemplo" />
</p>

O **Sistema de Cálculo de Rescisão Trabalhista** é uma aplicação web desenvolvida em **Go** (backend) e **PHP/HTML/CSS/JS** (frontend), com o framework **Fiber**, que expõe uma API RESTful. O sistema calcula as verbas rescisórias de um contrato de trabalho, incluindo salários proporcionais, décimo terceiro salário, férias e FGTS, além de realizar validações dos dados fornecidos.

### Funcionalidades

- **Cálculo das Verbas Rescisórias:**
  - Salário diário.
  - Décimo terceiro salário.
  - Férias proporcionais.
  - FGTS.
  
- **Validação dos Dados de Entrada:**
  - Validação das datas de admissão e demissão.
  - Validação do salário.
  - Validação do tipo de rescisão (Justa Causa, Sem Justa Causa ou Demissão).

### Arquitetura

- **Entidades:**
  - `EmploymentContract`: Representa o contrato de trabalho, com dados como data de admissão, data de demissão, salário e tipo de rescisão.
  - `ContractCalculationResults`: Contém os resultados dos cálculos de rescisão, como salários proporcionais, FGTS, férias e 13º salário.

- **Usecase:**
  - Lógica de negócios para calcular as verbas rescisórias.

- **Adapter:**
  - Camada de adaptação para expor a API RESTful.

- **Main:**
  - Função principal que inicia a aplicação.

### Tecnologias

- **Go:** Linguagem de programação.
- **Fiber:** Framework para desenvolvimento de APIs RESTful.
- **Testify:** Biblioteca para teste unitários.

---
## Como Usar

### 1. **Instale as Dependências**

Execute o comando abaixo para instalar as dependências do projeto:

```bash
go mod tidy

### 2. Execute o Servidor
   
Execute o comando abaixo para iniciar o servidor:

```bash
go run main.go

### 3. Envie uma Requisição POST para / com os dados do contrato de trabalho no corpo da requisição

### 4. Endpoints API:

POST /: Cria um novo cálculo de rescisão trabalhista, retornando os valores das verbas rescisórias calculadas.
```

---

## Alterações necessárias

- No arquivo ```contact.php```, altere ```"contato@dominio.com"``` para o endereço de email que deseja receber o conteúdo do formulário de contato
- Dentro da pasta ```tabs```, no arquivo ```index.php```, altere o parâmetro ```url``` para o endereço do servidor que trará o retorno
