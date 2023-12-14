import {
  Card,
  CardActions,
  CardContent,
  Button,
  Stepper,
  Step,
  StepLabel,
  TextField,
  Modal,
  Box,
} from "@mui/material";
import CardSkeleton from "./CardSkelton";
import { useState } from "react";
import CustomSnackBar, { Severity } from "@/components/SnackBar";
import Cards from "react-credit-cards-2";
import "react-credit-cards-2/dist/es/styles-compiled.css";

const style = {
  position: "absolute" as "absolute",
  top: "50%",
  left: "50%",
  transform: "translate(-50%, -50%)",
  width: 400,
  bgcolor: "background.paper",
  border: "2px solid #000",
  boxShadow: 24,
  pt: 2,
  px: 4,
  pb: 3,
};

type CreditCardProps = {
  onSubmit: (
    userId: string,
    cardNumber: string,
    expiryDate: string,
    cvv: string
  ) => void;
  onChangeCardNumber: (cardNumber: string) => void;
  userId: string;
  currentCardNumber: string;
};

const CreditCard = ({
  userId,
  currentCardNumber,
  onSubmit,
  onChangeCardNumber,
}: CreditCardProps) => {
  const [open, setOpen] = useState<boolean>(false);
  const [cardNumber, setCardNumber] = useState<string>("");
  const [expiryDate, setExpiryDate] = useState<string>("");
  const [cvv, setCVV] = useState<string>("");

  const handleOpen = () => setOpen(true);
  const handleClose = () => setOpen(false);
  const handleSubmit = () => {
    onSubmit(userId, cardNumber, expiryDate, cvv);
    onChangeCardNumber(cardNumber);

    handleClose();
  };
  const onChangeCardNumberHandler = (
    e: React.ChangeEvent<HTMLInputElement>
  ) => {
    const input = e.target.value.replace(/\D/g, "").substring(0, 16);
    const newInput = input.replace(/(\d{4})/g, "$1-");
    e.target.value = newInput.substring(0, 19);
    setCardNumber(input);
  };
  const onChangeExpiryDateHandler = (
    e: React.ChangeEvent<HTMLInputElement>
  ) => {
    const input = e.target.value.replace(/\D/g, "").substring(0, 4);
    const newInput = input.replace(/(\d{2})/g, "$1/");
    e.target.value = newInput.substring(0, 5);
    setExpiryDate(input);
  };
  const helperTextForCardNumber =
    cardNumber.replace(/\D/g, "").length !== 16
      ? "Please enter a valid 16-digit card number"
      : "";
  const helperTextForExpiryDate =
    expiryDate.replace(/\D/g, "").length !== 4
      ? "Please enter a valid 4-digit expiry date"
      : "";
  const helperTextForCVV =
    cvv.replace(/\D/g, "").length !== 3
      ? "Please enter a valid 3-digit CVV"
      : "";

  return (
    <>
      <Card sx={{ maxWidth: 700, minWidth: 100, minHeight: 50 }}>
        {currentCardNumber ? (
          <CardContent>****{currentCardNumber.slice(-4)}</CardContent>
        ) : (
          <CardSkeleton />
        )}
      </Card>
      <Button size="small" onClick={handleOpen}>
        Change Card
      </Button>

      <Modal
        open={open}
        onClose={handleClose}
        aria-labelledby="parent-modal-title"
        aria-describedby="parent-modal-description"
      >
        <Card
          sx={{
            ...style,
            maxWidth: 1500,
            width: "50%",
            Height: 900,
            display: "flex",
            justifyContent: "space-around",
          }}
        >
          <Box sx={{ display: "flex", alignItems: "center" }}>
            <Cards cvc={cvv} expiry={expiryDate} name="" number={cardNumber} />
          </Box>

          <Box sx={{ width: "300px" }}>
            <CardContent>
              Current Card: ****{currentCardNumber.slice(-4)}
            </CardContent>
            <CardContent>Please fill in new credit card</CardContent>
            <TextField
              fullWidth
              label="Card Number"
              onChange={onChangeCardNumberHandler}
              error={
                cardNumber ? cardNumber.replace(/\D/g, "").length !== 16 : false
              }
              helperText={helperTextForCardNumber}
            />
            <Box sx={{ display: "flex" }}>
              <TextField
                label="Expiry Date"
                error={
                  expiryDate
                    ? expiryDate.replace(/\D/g, "").length !== 4
                    : false
                }
                onChange={onChangeExpiryDateHandler}
                helperText={helperTextForExpiryDate}
              />
              <TextField
                label="CVV"
                error={cvv ? cvv.replace(/\D/g, "").length !== 3 : false}
                onChange={(e) => {
                  e.target.value = e.target.value
                    .replace(/\D/g, "")
                    .substring(0, 3);
                  setCVV(e.target.value);
                }}
                helperText={helperTextForCVV}
              />
            </Box>
            <CardActions>
              <Button
                color="success"
                onClick={handleSubmit}
                sx={{ marginLeft: "auto" }}
              >
                Submit
              </Button>
            </CardActions>
          </Box>
        </Card>
      </Modal>
    </>
  );
};

export default CreditCard;
