import { POSSIBLES_SYSTEMS } from "./../enums/enums";
import {
  getRandomBool,
  getRandomDate,
  getRandomIntInRange,
} from "../utils/rand";
import { OccurrenceDetailsInterface } from "./../interfaces/interfaces";

export function GenerateOccurrencesMockData(): OccurrenceDetailsInterface[] {
  let mockList: OccurrenceDetailsInterface[] = [];

  for (let i = 0; i <= getRandomIntInRange(0, 100); i++) {
    let tags = [];
    for (let index = 0; index < getRandomIntInRange(0, 4); index++) {
      tags.push(`TAG_${index}`);
    }

    mockList.push({
      AccountId: getRandomIntInRange(1000, 9999),
      OccurrenceId: getRandomIntInRange(1000, 9999),
      CreatedAt: getRandomDate(),
      System:
        POSSIBLES_SYSTEMS[getRandomIntInRange(0, POSSIBLES_SYSTEMS.length - 1)],
      Title: "Mock template data",
      Description: "This data were created using random functions",
      Image: "abcdef",
      Location: "-27.000, 32.9123",
      Tags: tags,
      AlreadyAnswered: getRandomBool(),
      AcceptedBy: "mock-user@gmail.com"
    });
  }

  return mockList;
}
