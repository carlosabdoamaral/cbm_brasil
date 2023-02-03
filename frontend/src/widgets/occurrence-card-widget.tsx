import { Card, Header, Icon } from "semantic-ui-react";
import { OccurrenceDetailsInterface } from "../interfaces/interfaces";
import { DATE_FORMATS } from "../enums/enums";
import Moment from "react-moment";
import Spacer from "./spacer-widget";
import { useEffect } from "react";

export default function OccurrenceCardWidget(occurrence: any) {
  useEffect(() => {
    console.log(occurrence);
  }, []);

  return (
    <Card fluid>
      <Card.Content>
        <Card.Header>#{occurrence.IdOccurrence.toString()}</Card.Header>

        <Card.Meta>
          <small className="date">
            <Moment format={DATE_FORMATS.DD_MM_YYYY}>
              {occurrence.CreatedAt}
            </Moment>
          </small>
        </Card.Meta>

        <Card.Description>
          <Spacer height={10} />
          <b>{occurrence.Title}</b>
          <br />
          <small>{occurrence.Description}</small>
        </Card.Description>
      </Card.Content>
      <Card.Content extra>
        <p color="red">
          {true && (
            <>
              <Icon name="check" />
              Ocorrência sendo atendida
            </>
          )}

          {!true && (
            <>
              <Icon name="close" />
              Ocorrência ainda não atendida
            </>
          )}
        </p>
      </Card.Content>
    </Card>
  );
}
