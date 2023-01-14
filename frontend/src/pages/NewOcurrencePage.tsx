import { useState } from "react";
import {
  Button,
  Container,
  Divider,
  Form,
  Input,
  TextArea,
} from "semantic-ui-react";
import { TitleWidget } from "../widgets/TitleWidget";

export function NewOcurrencePage() {
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const [image, setImage] = useState("");
  const [location, setLocation] = useState("");

  const renderFields = () => {
    return (
      <>
        <Form.Field>
          <label>Título</label>
          <Input
            placeholder="Título da ocorrência"
            value={title}
            onChange={(_, data) => {
              setTitle(data.value);
            }}
          />
        </Form.Field>

        <Form.Field>
          <label>Localização</label>
          <Input
            placeholder="Localização da ocorrencia, por padrão será usada sua localização"
            value={location}
            onChange={(_, data) => {
              setLocation(data.value);
            }}
          />
        </Form.Field>

        <Form.Field>
          <label>Imagem</label>
          <Input
            type="file"
            value={image}
            onChange={(_, data) => {
              setImage(data.value);
            }}
          />
        </Form.Field>

        <Form.Field>
          <label>Descrição</label>
          <TextArea
            style={{ minHeight: "200px" }}
            value={description}
            onChange={(_, data) => {
              setDescription(data.value!.toString());
            }}
            placeholder="Descreva rapidamente sua ocorrencia, isso irá ajudar os reponsáveis a tomar as medidas necessárias para ajudar"
          />
        </Form.Field>
      </>
    );
  };

  const renderSubmitButton = () => {
    return <Button type="submit" color="green" fluid content="Enviar" />;
  };

  //   const getGeoLocation = () => {
  //     if (!!navigator.geolocation) {
  //       navigator.geolocation.getCurrentPosition(showPosition);
  //     } else {
  //       console.log("Nao suporta GeoLocalizacao");
  //     }
  //   };

  //   const showPosition = (position: any) => {
  //     console.log(position.coords.latitude, position.coords.longitude);
  //   };

  return (
    <Container>
      <TitleWidget title="Nova ocorrência" subtitle="" />
      <Divider />

      <Form>
        {renderFields()}
        {renderSubmitButton()}
      </Form>
    </Container>
  );
}
