import { AddressInterface } from "./../interfaces/interfaces";
import axios from "axios";
import { addErrorNotification } from "./notifications";

export async function getCEPDetails(cep: string) {
  let addr: AddressInterface = {
    CEP: cep,
    Country: "Brazil",
    State: "EMPTY",
    City: "EMPTY",
    Neighborhood: "EMPTY",
    Street: "EMPTY",
    Ddd: "EMPTY",
    PlaceNumber: 0,
    PlaceNotes: "EMPTY",
  };

  await axios.get(`https://viacep.com.br/ws/${cep}/json/`).then((res) => {
    const resp = res.data;

    if (resp.erro) {
      addErrorNotification(
        "Não foi possível encontrar infomações do CEP fornecido"
      );
    } else {
      addr = {
        CEP: resp.cep,
        Country: "Brazil",
        State: resp.uf,
        City: resp.localidade,
        Neighborhood: resp.bairro,
        Street: resp.logradouro,
        Ddd: resp.ddd,
        PlaceNumber: 0,
        PlaceNotes: "",
      };
    }
  });

  return addr;
}
