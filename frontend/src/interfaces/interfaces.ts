export interface NewOccurrenceRequestInterface {
  AccountId: Number;
  CreatedAt: Date;
  System: String;
  Title: String;
  Description: String;
  Image: String;
  Location: String;
}

export interface OccurrenceDetailsInterface {
  AccountId: Number;
  OccurrenceId: Number;
  CreatedAt: Date;
  System: String;
  Title: String;
  Description: String;
  Image: String;
  Location: String;
  Tags: String[];
  AlreadyAnswered: Boolean;
  AcceptedBy: string;
}
