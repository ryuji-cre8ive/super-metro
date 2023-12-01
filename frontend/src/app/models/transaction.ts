export interface Transaction {
  id: string;
  userId: string;
  paymentId: string;
  amount: number;
  transactionType: string;
  createdAt: Date;
  updatedAt: Date;
  deletedAt: Date | null;
}
