package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"
)

const yaml = `
difficulty: {{.Difficulty}}
preparation_time: {{.PreparationTime}}
cooking_time: {{.CookingTime}}
food_yield: {{.FoodYield}}
ingredients: {{.Ingredients}}
method_of_preparation: {{.MethodOfPreparation}}
example_combinations: {{.ExampleCombinations}}
can_be_used_in_cooking: {{.CanBeUsedInCooking}}
with_chron: {{.WithChron}}
`

const html = `
<div align="justify" class="western">
<br />
<div class="separator" style="clear: both; text-align: center;">
</div>
<div class="separator" style="clear: both; text-align: center;">
</div>
<div class="separator" style="clear: both; text-align: center;">
</div>
<br />
Dificuldade: {{.Difficulty}}<br />
<br />
Tempo de preparo: {{.PreparationTime}}&nbsp; Tempo de cozimento: {{.CookingTime}}<br />
Rendimento: {{.FoodYield}}<br />
<!--more-->
<br />
<ul>
{{ range $key, $value := .Ingredients }}
<li>{{$value}}</li>
{{ end }}
</ul>
<br />
{{ range $key, $value := .MethodOfPreparation }}
{{ $value }}<br />
{{ end }}
<br />
Exemplo para combinações: {{.ExampleCombinations}}<br />
<br />
{{if .CanBeUsedInCooking}}
Essa receita, pode ser usada (com algumas restrições) para caldos (geralmente em preparo de alguns exames)<br />
<br />
{{end}}
{{if .WithChron}}
<b>Essa receita deve ser usada com moderação para pacientes com Crohn. Uso não recomendado para quem está com o Crohn em atividade.</b>
<br />
{{end}}
<br />
*A ideia de cardápio pode variar e deve ser tratado de acordo com seus hábitos e preferências.<br />
Caso haja alguma restrição médica, sugiro ler <a href="http://dietasimsgn.blogspot.com.br/search/label/Propriedades%20Nutricionais">Propriedades Nutricionais</a> e <a href="http://dietasimsgn.blogspot.com.br/p/blog-page.html">tabela de substituições</a> dos alimentos.</div>
`

const EOL = '\n'

type Blogging struct {
	Filename,
	Difficulty,
	PreparationTime,
	CookingTime,
	FoodYield,
	ExampleCombinations string
	Ingredients,
	MethodOfPreparation []string
	CanBeUsedInCooking,
	WithChron bool
}

func checkIfTrue(field string) bool {
	if field == "true" {
		return true
	}
	return false
}

func splitSemicolon(str string) (text []string) {
	ss := strings.Split(str, "\r\n")[0]
	for _, s := range strings.Split(ss, ";") {
		var strs []string
		for _, st := range strings.Split(s, " ") {
			if st != "" {
				strs = append(strs, st)
			}
		}
		text = append(text, strings.Join(strs, " "))
	}
	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var blogging Blogging

	fmt.Println("Título do arquivo: ")
	filename, _ := reader.ReadString(EOL)
	yamlFilename := strings.Split(filename, "\r\n")[0]
	yamlF := []string{yamlFilename, ".yml"}
	blogging.Filename = strings.Join(yamlF, "")

	fmt.Println("Dificuldade: ")
	difficulty, _ := reader.ReadString(EOL)
	blogging.Difficulty = strings.Split(difficulty, "\r\n")[0]

	fmt.Println("Tempo de preparo: ")
	preparationTime, _ := reader.ReadString(EOL)
	blogging.PreparationTime = strings.Split(preparationTime, "\r\n")[0]

	fmt.Println("Tempo de cozimento: ")
	cookingTime, _ := reader.ReadString(EOL)
	blogging.CookingTime = strings.Split(cookingTime, "\r\n")[0]

	fmt.Println("Rendimento: ")
	foodYield, _ := reader.ReadString(EOL)
	blogging.FoodYield = strings.Split(foodYield, "\r\n")[0]

	fmt.Println("Ingredientes: ")
	ingredients, _ := reader.ReadString(EOL)
	blogging.Ingredients = append(blogging.Ingredients, splitSemicolon(ingredients)...)

	fmt.Println("Modo de preparo: ")
	methodOfPreparation, _ := reader.ReadString(EOL)
	blogging.MethodOfPreparation = append(blogging.MethodOfPreparation, splitSemicolon(methodOfPreparation)...)

	fmt.Println("Exemplo para combinações: ")
	exampleCombinations, _ := reader.ReadString(EOL)
	blogging.ExampleCombinations = strings.Split(exampleCombinations, "\r\n")[0]

	fmt.Println("Pode ser usado no preparo?(true ou false): ")
	canBeUsedInCooking, _ := reader.ReadString(EOL)
	blogging.CanBeUsedInCooking = checkIfTrue(strings.Split(canBeUsedInCooking, "\r\n")[0])

	fmt.Println("Para pacientes com chron?(true ou false): ")
	withChron, _ := reader.ReadString(EOL)
	blogging.WithChron = checkIfTrue(strings.Split(withChron, "\r\n")[0])

	tYaml := template.Must(template.New("yaml").Parse(yaml))
	outputYaml, err := os.Create(blogging.Filename)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := outputYaml.Close(); err != nil {
			panic(err)
		}
	}()

	if err := tYaml.Execute(outputYaml, blogging); err != nil {
		panic(err)
	}

	htmlFilename := strings.Split(filename, "\r\n")[0]
	htmlF := []string{htmlFilename, ".html"}
	blogging.Filename = strings.Join(htmlF, "")

	tHtml := template.Must(template.New("html").Parse(html))
	outputHml, err := os.Create(blogging.Filename)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := outputHml.Close(); err != nil {
			panic(err)
		}
	}()

	if err := tHtml.Execute(outputHml, blogging); err != nil {
		panic(err)
	}
}
