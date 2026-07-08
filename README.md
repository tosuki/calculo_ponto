# Calculadora de Saída

Uma aplicação simples em Python para calcular o horário de saída do trabalho com base na hora de entrada, tempo de almoço e jornada de trabalho.

## Funcionalidades

- Entrada de horário de início no formato `HH:MM`
- Registro do tempo de almoço em minutos
- Definição da jornada de trabalho em horas
- Exibe o horário previsto de saída
- Atualização em tempo real da hora atual e do tempo restante
- Indica quando você está em hora extra

## Tecnologias

- Python 3
- customtkinter
- datetime

## Como executar

1. Instale Python 3 caso ainda não tenha.
2. Instale a dependência:

```bash
pip install customtkinter
```

3. Execute o aplicativo:

```bash
python app.py
```

## Observações

- O aplicativo usa `customtkinter` para interface gráfica.
- A janela tem tamanho fixo de `420x420`.
- Caso os dados de entrada estejam inválidos, a aplicação exibirá "Dados inválidos.".

## Estrutura do projeto

- `app.py` - código principal da aplicação
- `app.spec` - especificação para empacotamento com PyInstaller
- `build/` - saída de build gerada pelo PyInstaller

## Uso

1. Informe a hora de entrada em `HH:MM`.
2. Informe o tempo de almoço em minutos.
3. Informe a jornada de trabalho em horas.
4. A tela exibirá o horário de saída e o tempo restante.

---

Bom trabalho e feliz cálculo de ponto!