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
`

const html = `
<div align="justify" class="western">
<div class="separator" style="clear: both; text-align: center;">
<a href="http://4.bp.blogspot.com/-qTF7Afr_VlM/VM36cKt0vmI/AAAAAAAAAPE/kErLM33-Yg4/s1600/paes%2Brecheados%2Bvertical.jpg" imageanchor="1" style="margin-left: 1em; margin-right: 1em;"><img border="0" src="http://4.bp.blogspot.com/-qTF7Afr_VlM/VM36cKt0vmI/AAAAAAAAAPE/kErLM33-Yg4/s1600/paes%2Brecheados%2Bvertical.jpg" height="320" width="305" /></a></div>
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
<br />
<ul>
<li>1 xícara de chá de água</li>
<li>1 ovo </li>
<li>1 gema (para pincelar)</li>
<li>3 colheres de sopa de açúcar</li>
<li>2 colheres de sopa de azeite para cozinhar (não deve ser extra virgem)</li>
<li>30 g ou 2 tabletes de fermento biológico ou de padaria</li>
<li>1 colher de chá de sal</li>
<li>2 colheres de sopa de gergelim</li>
<li>2 xícaras de farinha de trigo integral (aproximadamente)</li>
<li>2 xícaras farinha de trigo branca</li>
<li>(pode fazer todo o processo com uma única farinha, eu uso 2 pois prefiro fazer um integral e uma comum)</li>
</ul>
<br />
Bata no mix ou liquidificador, todos os ingredientes exceto as farinhas e o gergelim (Passo 3)<br />
Despeje METADE do líquido em uma vasilha, junte o gergelim e a farinha de trigo integral aos poucos e amasse bem, até que a massa fique lisa e homogênea. (passo 4)<br />
Deixe descansar por 30 minutos<br />
Aqueça o forno com temperatura de 200ºC<br />
Despeje o restante do líquido em uma vasilha, junte a farinha de trigo branca aos poucos e amasse bem, até que a massa fique lisa e homogênea. (Passo 5)<br />
Deixe descansar por 30 minutos<br />
Abra a primeira massa com o auxílio de um rolo em superfície enfarinhada, recheie, enrole o pão e pincele 1 gema (passo 6)<br />
Abra a segunda massa com o auxílio de um rolo em superfície enfarinhada, recheie, enrole o pão e pincele 1 gema. (passo 10)<br />
coloque as massas em uma assadeira (nesse caso, cortei em fatias e coloquei na forma de cupcakes para ficar rústico) e asse por aproximadamente 30 minutos.<br />
<br />
Exemplo para combinações: {{.ExampleCombinations}}<br />
<br />
<div class="separator" style="clear: both; text-align: center;">
<a href="http://4.bp.blogspot.com/-rusL2wh86mk/VM2SSF9GJLI/AAAAAAAAAO0/E3oZeWvjUsU/s1600/pao%2Brecheado%2Bfinal.jpg" imageanchor="1" style="margin-left: 1em; margin-right: 1em;"><img border="0" src="http://4.bp.blogspot.com/-rusL2wh86mk/VM2SSF9GJLI/AAAAAAAAAO0/E3oZeWvjUsU/s1600/pao%2Brecheado%2Bfinal.jpg" height="120" width="320" /></a></div>
<br />
<br />
*A ideia de cardápio pode variar e deve ser tratado de acordo com seus hábitos e preferências.<br />
Caso haja alguma restrição médica, sugiro ler <a href="http://dietasimsgn.blogspot.com.br/search/label/Propriedades%20Nutricionais">Propriedades Nutricionais</a> e <a href="http://dietasimsgn.blogspot.com.br/p/blog-page.html">tabela de substituições</a> dos alimentos.</div>
`

type Blogging struct {
	Filename,
	Difficulty,
	PreparationTime,
	CookingTime,
	FoodYield,
	Ingredients,
	MethodOfPreparation,
	ExampleCombinations,
	CanBeUsedInCooking string
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var blogging Blogging

	fmt.Println("Título do arquivo: ")
	filename, _ := reader.ReadString('\n')
	yamlFilename := strings.Split(filename, "\n")[0]
	yamlF := []string{yamlFilename, ".yml"}
	blogging.Filename = strings.Join(yamlF, "")

	fmt.Println("Dificuldade: ")
	difficulty, _ := reader.ReadString('\n')
	blogging.Difficulty = strings.Split(difficulty, "\n")[0]

	fmt.Println("Tempo de preparo: ")
	preparationTime, _ := reader.ReadString('\n')
	blogging.PreparationTime = strings.Split(preparationTime, "\n")[0]

	fmt.Println("Tempo de cozimento: ")
	cookingTime, _ := reader.ReadString('\n')
	blogging.CookingTime = strings.Split(cookingTime, "\n")[0]

	fmt.Println("Rendimento: ")
	foodYield, _ := reader.ReadString('\n')
	blogging.FoodYield = strings.Split(foodYield, "\n")[0]

	fmt.Println("Ingredientes: ")
	ingredients, _ := reader.ReadString('\n')
	blogging.Ingredients = strings.Split(ingredients, "\n")[0]

	fmt.Println("Modo de preparo: ")
	methodOfPreparation, _ := reader.ReadString('\n')
	blogging.MethodOfPreparation = strings.Split(methodOfPreparation, "\n")[0]

	fmt.Println("Exemplo para combinações: ")
	exampleCombinations, _ := reader.ReadString('\n')
	blogging.ExampleCombinations = strings.Split(exampleCombinations, "\n")[0]

	fmt.Println("Pode ser usado no preparo?: ")
	canBeUsedInCooking, _ := reader.ReadString('\n')
	blogging.CanBeUsedInCooking = strings.Split(canBeUsedInCooking, "\n")[0]

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

	htmlFilename := strings.Split(filename, "\n")[0]
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
