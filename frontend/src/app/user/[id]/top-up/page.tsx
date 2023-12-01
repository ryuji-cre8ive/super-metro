"use client";
import React, { useState, useEffect } from "react";
import { TextField, Typography } from "@mui/material";
import LoadingButton from "@mui/lab/LoadingButton";
import SendIcon from "@mui/icons-material/Send";
import Image from "next/image";
import CustomSnackBar, { Severity } from "@/components/SnackBar";
import EWalletPic from "../../../../../public/e-wallet.png";
import {
  SnackBarTextErr,
  SnackBarTextSuccess,
  SnackBarTextInputErr,
  SnackBarTextInputSuccess,
} from "./text";
import { useAuth } from "@/app/userContext";
import axios from "@/api/axiosConfig";
import { AxiosResponse } from "axios";
import Card from "./components/Card";

const TopUp = () => {
  const { user, topUp } = useAuth();
  const [amount, setAmount] = useState<number | null>(null);
  const [errOpen, setErrOpen] = useState<boolean>(false);
  const [successOpen, setSuccessOpen] = useState<boolean>(false);
  const [text, setText] = useState<string>("");
  const [loading, setLoading] = useState<boolean>(false);
  const [cardNumber, setCardNumber] = useState<string>("");

  useEffect(() => {
    if (!user) return;
    getCardNumber(user.id);
  }, [user]);

  const getCardNumber = async (userId: string) => {
    axios.get(`/credit-card/${userId}`).then((res) => {
      if (!res.data) {
        return;
      }
      res.data.CardNumber && setCardNumber(res.data.CardNumber);
    });
  };

  const onChangeCardNumber = (cardNumber: string) => {
    setCardNumber(cardNumber);
  };

  const onSubmit = async (
    userId: string,
    cardNumber: string,
    expiryDate: string,
    cvv: string
  ) => {
    const params = {
      userId,
      cardNumber,
      expiryDate,
      cvv,
    };
    if (!userId || !cardNumber || !expiryDate || !cvv) {
      setText(SnackBarTextInputErr);
      return setErrOpen(true);
    }
    try {
      await axios.post("/credit-card/add", params).then((res) => {
        if (res.status !== 200) {
          setText(SnackBarTextInputErr);
          setErrOpen(true);
        }
        setText(SnackBarTextInputSuccess);
        setSuccessOpen(true);
      });
    } catch (e) {
      setText(SnackBarTextInputErr);
      setErrOpen(true);
    }
  };

  const isNumber = (value: any) => {
    var pattern = /^[0-9]*$/;
    return pattern.test(value);
  };

  const handleAmount = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    if (!isNumber(e.target.value)) {
      return setErrOpen(true);
    }
    setErrOpen(false);
    setAmount(Number(e.target.value));
  };

  const handleTopUp = () => {
    if (!amount) return;
    if (!isNumber(amount)) {
      setText(SnackBarTextErr);
      return setErrOpen(true);
    }
    if (!user) return;
    setLoading(true);
    try {
      axios
        .post("/top-up", { valance: Number(amount), id: user.id })
        .then((res: AxiosResponse) => {
          if (res.status !== 200) return;
          topUp(user, Number(amount));
          window.localStorage.setItem("session_token", res.data.sessionToken);
          setText(SnackBarTextSuccess);
          setSuccessOpen(true);
          setAmount(null);
          setLoading(false);
        });
    } catch (e) {
      setLoading(false);
    }
  };

  return (
    <main
      style={{
        margin: "30px",
        textAlign: "center",
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
      }}
    >
      <Image src={EWalletPic} alt="e-wallets" loading="lazy" width={300} />
      <Card
        onSubmit={onSubmit}
        currentCardNumber={cardNumber}
        userId={user ? user.id : ""}
        onChangeCardNumber={onChangeCardNumber}
      />
      {}
      <Typography variant="subtitle1">
        Valance: {user && user.valance}
      </Typography>
      <TextField label="Amount" onChange={handleAmount} />
      <LoadingButton
        color="primary"
        onClick={handleTopUp}
        loading={loading}
        loadingPosition="center"
        endIcon={<SendIcon />}
        variant="contained"
      >
        <span>TopUp!!</span>
      </LoadingButton>
      <CustomSnackBar
        open={errOpen}
        setOpen={setErrOpen}
        severity={Severity.ERROR}
        text={text}
      />
      <CustomSnackBar
        open={successOpen}
        setOpen={setSuccessOpen}
        severity={Severity.SUCCESS}
        text={text}
      />
    </main>
  );
};

export default TopUp;
