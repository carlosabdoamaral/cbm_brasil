import { MOCK_TEMPLATES } from "./../enums/enums";
import { GenerateOccurrencesMockData } from "./occurrences";

export function GenerateMock(template: string = MOCK_TEMPLATES.OCCURRENCES) {
  if (template === MOCK_TEMPLATES.OCCURRENCES) {
    return GenerateOccurrencesMockData();
  }
}
