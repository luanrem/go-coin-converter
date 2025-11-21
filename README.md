# Conversor de Moedas (BRL para Outras Moedas) - CLI

## Visão Geral

Este projeto solicita que você construa uma ferramenta de linha de comando em Go para converter valores em Reais para outras moedas utilizando taxas de câmbio pré-definidas.

## Requisitos Funcionais

1. **Interface de Linha de Comando:** O programa deve ser executável via terminal.
2. **Argumentos de Entrada:** A aplicação deve aceitar exatamente dois argumentos na seguinte ordem:
   1. `[valor_em_brl]`: Um número (float) representando a quantia em Reais a ser convertida.
   2. `[moeda_destino]`: O código da moeda para a qual o valor será convertido (ex: `USD`, `EUR`, `JPY`).
3. **Fonte das Taxas:** O programa deve utilizar as taxas de câmbio fornecidas abaixo. Você pode optar por:
   - Carregar os dados de um arquivo `rates.json` local.
   - Incorporar (hardcode) os dados diretamente no seu código Go (por exemplo, usando um `map`).
4. **Cálculo:** O programa deve multiplicar o `[valor_em_brl]` pela taxa correspondente à `[moeda_destino]`.
5. **Saída:** A aplicação deve imprimir no console *apenas* o valor final convertido, formatado com duas casas decimais.

## Sintaxe de Comando Esperada

```bash
./convert [valor_em_brl] [moeda_destino]
```

## Exemplos de Execução

```bash
# Exemplo 1: Converter 10 BRL para USD
$ ./convert 10 USD
1.51

# Exemplo 2: Converter 12 BRL para JPY
$ ./convert 12 JPY
195.48
```

## Fonte de Dados (Taxas de Câmbio Base: BRL)

Utilize o JSON abaixo como fonte para as taxas de conversão.

```json
{
  "base": "BRL",
  "date": "2025-04-14",
  "rates": {
    "USD": 0.151,
    "EUR": 0.137,
    "JPY": 16.29,
    "GBP": 0.13,
    "CHF": 0.1402,
    "AUD": 0.2712
  }
}
```

## Recursos Úteis (Referências)

- **Argumentos de Linha de Comando:** https://gobyexample.com/command-line-arguments
- **Leitura de Arquivos (se optar por essa via):** https://gobyexample.com/reading-files
- **Processamento de JSON (Unmarshal):** https://gobyexample.com/json
- **Uso de Maps (se optar por hardcode):** https://gobyexample.com/maps

## Conceitos Praticados no Projeto

Ao desenvolver este projeto, você praticará os seguintes conceitos de programação em Go:

- **Argumentos da Linha de Comando:** Ler dados que o usuário passa pelo terminal (`os.Args`).
- **Estruturas de Dados:** Usar `maps` para armazenar e acessar as taxas de câmbio.
- **Conversão de Tipos:** Transformar texto (string) em um número (`float`) para realizar cálculos.
- **Manipulação de Strings:** Modificar textos, como converter a moeda para letras maiúsculas.
- **Lógica Condicional:** Verificar as entradas do usuário com `if/else`.
- **Tratamento de Erros:** Lidar com entradas inválidas (número errado de argumentos, moeda inexistente, etc.).
- **Formatação de Saída:** Imprimir o resultado final com um formato específico (ex: duas casas decimais).

## Desafios Extras (Opcional)

1. **Tratamento de Erros:** Melhore o programa para lidar com entradas inválidas:
   - Número incorreto de argumentos (ex: `./convert 10`).
   - Valor em BRL que não é um número (ex: `./convert abc USD`).
   - Código de moeda que não existe nas taxas (ex: `./convert 10 XYZ`).
2. **Formatação de Saída:** Garanta que a saída numérica respeite a formatação local para decimais (vírgula em vez de ponto), se aplicável. *(Nota: O exemplo `195,48` sugere isso, mas o exemplo `1.50` sugere o contrário. O ideal é padronizar - para CLI, o ponto `.` é mais comum e fácil de usar em scripts)*.
3. **Sensibilidade a Maiúsculas/Minúsculas:** Faça o programa funcionar independentemente se o usuário digitar `usd` ou `USD`.
