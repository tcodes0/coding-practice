#! /usr/bin/env node

/**
 * Nessa funcao um array de pessoas e um array de comidas sao criados
 * A saida desejada:
 * [
 * "jac come pizza",
 * "tom come sorvete",
 * "bel come pizza",
 * "andrea come sorvete",
 * "bruno come pizza",
 * ]
 *
 * A saida atual:
 * * [
 * <1 empty item>,
 * "jac  come sorvete",
 * "tom  come undefined",
 * "bel  come undefined",
 * "andrea  come undefined",
 * "bruno  come undefined",
 * ]
 *
 * Corrija a saida do programa
 * 1 - remove o item estranho "empty", deve ser "jac come sorvete"
 * 2 - remova undefined, "tom come undefined" deve ser "tom come pizza"
 * 3 - alterne entre pizza e sorvete
 * 4 - remova os espacos em branco extra entre o nome e a comida
 *
 * Uma vez corrigido
 * 1- reescreva a funcao para ser mais simples
 */
function pizzaSorvete() {
  let people = ["jac", "tom", "bel", "andrea", "bruno"];
  let foods = ["pizza", "sorvete"];

  let i = 1;
  let max = people.length;
  let out = [];

  for (let p of people) {
    out[i] = p + " " + " come " + foods[i];

    if (i === max) {
      break;
    }

    i++;
  }

  console.log(out);
}

try {
  pizzaSorvete();
} catch (e) {
  console.error("error: " + e);
}
