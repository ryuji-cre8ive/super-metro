"use client";
import {
  Timeline,
  TimelineItem,
  TimelineOppositeContent,
  TimelineSeparator,
  TimelineDot,
  TimelineConnector,
  TimelineContent,
  timelineOppositeContentClasses,
} from "@mui/lab";
import {
  Container,
  Grid,
  Card,
  CardHeader,
  CardContent,
  Link,
  Typography,
  Alert,
} from "@mui/material";
import { Transaction } from "@/app/models/transaction";

type TransactionProps = {
  transactions: Transaction[];
};

const RecentTransactions = ({ transactions }: TransactionProps) => {
  const getColor = (transactionType: string) => {
    switch (transactionType) {
      case "TOPUP":
        return "success";
      case "USE":
        return "secondary";
      default:
        return "warning";
    }
  };
  const getDescription = (transactionType: string) => {
    switch (transactionType) {
      case "TOPUP":
        return "top-up with ";
      case "USE":
        return "use your wallet to pay for ";
      default:
        return "不明なトランザクションタイプ";
    }
  };
  const convertToLocalTime = (time: Date) => {
    const date = new Date(time);
    return date.toLocaleString();
  };

  return (
    <Container maxWidth="lg">
      <Grid container justifyContent="center" mt={3} spacing={3}>
        <Grid item xs={12} lg={4}>
          <Card sx={{ paddingBottom: "30px" }}>
            <CardHeader
              title="Recent Transactions"
              titleTypographyProps={{ fontSize: "18px" }}
            />

            <CardContent>
              <Timeline
                className="theme-timeline"
                nonce={undefined}
                onResize={undefined}
                onResizeCapture={undefined}
                sx={{
                  p: 0,
                  mb: "-40px",
                  [`& .${timelineOppositeContentClasses.root}`]: {
                    flex: 0.5,
                    paddingLeft: 0,
                  },
                }}
              >
                {transactions.map((transaction, index) => (
                  <TimelineItem key={index}>
                    <TimelineOppositeContent>
                      {convertToLocalTime(transaction.createdAt)}
                    </TimelineOppositeContent>
                    <TimelineSeparator>
                      <TimelineDot
                        color={getColor(transaction.transactionType)}
                        variant="outlined"
                      />
                      {transactions.length - 1 != index && (
                        <TimelineConnector />
                      )}
                    </TimelineSeparator>
                    <TimelineContent>
                      {getDescription(transaction.transactionType) +
                        transaction.amount}
                    </TimelineContent>
                  </TimelineItem>
                ))}
              </Timeline>
            </CardContent>
          </Card>
        </Grid>
      </Grid>
    </Container>
  );
};

export default RecentTransactions;
