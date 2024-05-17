export interface User {
  firstName?: string,
  lastName?: string,
  middleInitial?: string,
  userName?: string,
  email?: string,
  phone?: string,
  password?: string,
  confirmPassword?: string,
  createdAt?: Date,
  lastUpdated?: Date,
  securityGroups?: string[]
}
