import { useState } from "react";
import {
  Button,
  Container,
  Dimmer,
  Form,
  Grid,
  Header,
  Loader,
  Select,
} from "semantic-ui-react";
import { NewOccurrenceRequestInterface } from "../interfaces/interfaces";
import { doPost } from "../utils/http";
import {
  addErrorNotification,
  addSuccessNotification,
  addWarningNotification,
} from "../utils/notifications";
import { GetOperationalSystem } from "../utils/operational-system";
import Spacer from "../widgets/SpacerWidget";

export function NewOcurrencePage() {
  const categoryOptions = [
    { key: "FIRE", text: "Incêndio", value: "FIRE" },
    { key: "CAR_FIRE", text: "Carro em chamas", value: "CAR_FIRE" },
    { key: "ANIMALS", text: "Animal em perigo", value: "ANIMALS" },
  ];

  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const [image, setImage] = useState("");
  const [location, setLocation] = useState("");
  const [category, setCategory] = useState(Object);
  const [isLoading, setIsLoading] = useState(false);

  function allRequiredFieldsAreFilled(): boolean {
    return !!title && !!description && !!image && !!location && !!category;
  }

  function handleSubmit() {
    if (allRequiredFieldsAreFilled()) {
      setIsLoading(true);
      const body: NewOccurrenceRequestInterface = {
        AccountId: 1,
        CreatedAt: new Date(),
        System: GetOperationalSystem(),
        Title: title,
        Description: description,
        Image: image,
        Location: location,
      };

      doPost("/api/", body)
        .then((res) => {
          addSuccessNotification("Enviado com sucesso");
          handleClearValues();
        })
        .catch((err) => {
          addErrorNotification(err);
        })
        .finally(() => {
          setIsLoading(false);
        });
    } else {
      addErrorNotification("Todos os campos devem estar preenchidos");
    }
  }

  function handleClearValues(mustNotify: boolean = false) {
    setTitle("");
    setDescription("");
    setImage("");
    setLocation("");
    setCategory({ key: "", text: "", value: "" });

    if (mustNotify) {
      addSuccessNotification("Todos os campos foram limpos");
    }
  }

  function renderHeader() {
    return (
      <Header
        as="h1"
        style={{ fontSize: 35, fontFamily: "poppins", fontWeight: 500 }}
      >
        Crie sua ocorrência <br /> agora mesmo 🧯
      </Header>
    );
  }

  function renderForm() {
    return (
      <Form>
        <Grid columns={2}>
          <Grid.Row>
            <Grid.Column>
              <Form.Input
                fluid
                label="Título"
                placeholder="Título"
                value={title}
                onChange={(_, data) => {
                  setTitle(data.value);
                }}
              />
            </Grid.Column>
            <Grid.Column>
              <Form.Input
                fluid
                label="Localização"
                placeholder="Localização"
                value={location}
                onChange={(_, data) => {
                  setLocation(data.value);
                }}
              />
            </Grid.Column>
          </Grid.Row>

          <Grid.Row>
            <Grid.Column>
              <Form.Select
                label="Categoria"
                onChange={(_, data) => {
                  const category = categoryOptions.filter(
                    (category) => category.value === data.value
                  );

                  setCategory(category);
                }}
                options={categoryOptions}
                placeholder="Categoria"
                value={category.value}
              />
            </Grid.Column>
            <Grid.Column>
              <Form.Input
                type="file"
                fluid
                label="Imagem"
                placeholder="Imagem"
                onChange={(_, data) => {
                  setImage(data.value);
                }}
              />
            </Grid.Column>
          </Grid.Row>
        </Grid>

        <Spacer height={20} />

        <Form.TextArea
          fluid
          label="About"
          placeholder="Tell us more about you..."
          onChange={(_, data) => {
            setDescription(data.value!.toString());
          }}
        />

        <Spacer height={10} />

        <Grid columns={2}>
          <Grid.Row>
            <Grid.Column>
              <Button
                fluid
                content="Limpar valores"
                color="red"
                icon={"close"}
                onClick={() => {
                  handleClearValues(true);
                }}
              />
            </Grid.Column>

            <Grid.Column>
              <Button
                fluid
                content="Enviar"
                secondary
                icon={"send"}
                onClick={() => {
                  handleSubmit();
                }}
              />
            </Grid.Column>
          </Grid.Row>
        </Grid>
      </Form>
    );
  }

  return (
    <>
      <Dimmer active={isLoading}>
        <Loader active={isLoading} />
      </Dimmer>
      <Container
        style={{
          display: "flex",
          flexDirection: "column",
          height: "100vh",
          justifyContent: "center",
        }}
      >
        <Container>
          {renderHeader()}
          <Container style={{ height: "30px" }} />
          {renderForm()}
        </Container>
      </Container>
    </>
  );
}
