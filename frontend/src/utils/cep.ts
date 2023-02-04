import { AddressInterface } from "./../interfaces/interfaces";
import axios from "axios";
import { addErrorNotification } from "./notifications";

export const emtpyAddr: AddressInterface = {
  CEP: "EMPTY",
  Country: "EMPTY",
  State: "EMPTY",
  City: "EMPTY",
  Neighborhood: "EMPTY",
  Street: "EMPTY",
  Ddd: "EMPTY",
  PlaceNumber: 0,
  PlaceNotes: "EMPTY",
};

export async function getCEPDetails(cep: string) {
  let addr: AddressInterface = emtpyAddr;

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

export function addressInterfaceIsEmpty(addr: AddressInterface): boolean {
  return addr === emtpyAddr;
}

export function formatCEP(inputValue: string): string {
  let res: string = inputValue;

  let re = /^([\d]{5})\.*-*([\d]{3})/; // Pode usar ? no lugar do *
  if (re.test(inputValue)) {
    res = inputValue.replace(re, "$1-$2");
  }

  return res;
}
