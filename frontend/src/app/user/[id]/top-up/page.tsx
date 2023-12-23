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
import { Topup } from "@/components/TopUp";
import { useRouter } from "next/navigation";

const TopUp = () => {
  const { user, topUp } = useAuth();
  const [amount, setAmount] = useState<number | null>(null);
  const [errOpen, setErrOpen] = useState<boolean>(false);
  const [successOpen, setSuccessOpen] = useState<boolean>(false);
  const [text, setText] = useState<string>("");
  const [loading, setLoading] = useState<boolean>(false);
  const [cardNumber, setCardNumber] = useState<string>("");

  const router = useRouter();
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

  const onChangeAmount = (amount: number) => {
    setErrOpen(false);
    setAmount(amount);
  };

  const handleAmount = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    e.preventDefault();

    if (!e.target.value) {
      setAmount(null);
      return;
    }
    onChangeAmount(Number(e.target.value));
  };

  const handleTopUp = () => {
    if (!amount) return;
    if (!isNumber(amount)) {
      setText(SnackBarTextErr);
      return setErrOpen(true);
    }
    if (!user) return;
    setLoading(true);

    axios
      .post("/top-up", { valance: Number(amount), id: user.id })
      .then((res: AxiosResponse) => {
        if (res.status !== 200) {
          throw new Error("Request failed with status code " + res.status);
        }
        topUp(user, Number(amount));
        setText(SnackBarTextSuccess);
        setSuccessOpen(true);
        setAmount(null);
        setLoading(false);
      })
      .catch((e) => {
        setErrOpen(true);
        setLoading(false);
        if (e.response.status === 401) {
          setText("Please login first");
          return router.push("/signin");
        }
      });
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
      <Card
        onSubmit={onSubmit}
        currentCardNumber={cardNumber}
        userId={user ? user.id : ""}
        onChangeCardNumber={onChangeCardNumber}
      />
      <Topup
        currentBalance={user ? user.valance : 0}
        amount={amount}
        handleTopUp={handleTopUp}
        handleAmount={handleAmount}
        loading={loading}
      />
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
