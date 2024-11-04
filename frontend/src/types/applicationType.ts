export type ApplicationType = {
  ID: string;
  Title: string;
  Url: string;
  Company: string;
  Location: string;
  Salary: number;
  Description: string;
  Applied: boolean;
  FollowUp: boolean;
  Response: boolean;
};

export interface ApplicationDisplayProps {
  applications: ApplicationType[];
}
