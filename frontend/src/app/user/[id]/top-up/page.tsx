"use client";
import React, { useState, useEffect } from "react";
import CustomSnackBar, { Severity } from "@/components/SnackBar";
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

const TopUp = () => {
  const { user } = useAuth();
  const [amount, setAmount] = useState<number | null>(null);
  const [currentAmount, setCurrentAmount] = useState<number>(0); // [TODO
  const [errOpen, setErrOpen] = useState<boolean>(false);
  const [successOpen, setSuccessOpen] = useState<boolean>(false);
  const [text, setText] = useState<string>("");
  const [loading, setLoading] = useState<boolean>(false);
  const [cardNumber, setCardNumber] = useState<string>("");
  useEffect(() => {
    if (!user) return;
    getCardNumber(user.id);
    getAmount(user.id);
  }, [user]);

  const getCardNumber = async (userId: string) => {
    axios.get(`/credit-card/${userId}`).then((res) => {
      if (!res.data) {
        return;
      }
      res.data.CardNumber && setCardNumber(res.data.CardNumber);
    });
  };

  const getAmount = async (userId: string) => {
    await axios.get(`/amount/${userId}`).then((res) => {
      if (!res.data) return;
      console.log(res.data);
      res.data && setCurrentAmount(res.data);
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
        setText(SnackBarTextSuccess);
        setSuccessOpen(true);
        setAmount(null);
        setCurrentAmount(currentAmount + Number(amount));
        setLoading(false);
      });
    // .catch((e: AxiosError) => {
    //   setErrOpen(true);
    //   setLoading(false);
    //   if (e.status === 401) {
    //     setText("Please login first");
    //     return router.push("/signin");
    //   }
    // });
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
        currentBalance={currentAmount}
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
