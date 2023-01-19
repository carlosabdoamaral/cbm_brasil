export interface NewOccurrenceRequestInterface {
  AccountId: Number;
  CreatedAt: Date;
  System: string;
  Title: string;
  Description: string;
  Image: string;
  Location: string;
}

export interface OccurrenceDetailsInterface {
  AccountId: Number;
  OccurrenceId: Number;
  CreatedAt: Date;
  System: string;
  Title: string;
  Description: string;
  Image: string;
  Location: string;
  Tags: string[];
  AlreadyAnswered: Boolean;
  AcceptedBy: string;
}

export interface SignInInterface {
  Cpf: string;
  Password: string;
}

export interface AddressInterface {
  CEP: string;
  Country: string;      // By default it's Brazil - Disabled
  State: string;        // Will be filled by an API
  City: string;         // Will be filled by an API
  Neighborhood: string; // Will be filled by an API
  Street: string;       // Will be filled by an API
  Ddd: string;          // Will be filled by an API
  PlaceNumber: number;
  PlaceNotes: string;
}

export interface SignUpInterface {
  Cpf: string;
  FullName: string;
  Age: string;
  MainAddress: AddressInterface;
  Password: string;
}
