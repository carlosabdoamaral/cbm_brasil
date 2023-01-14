import { Store } from "react-notifications-component";

export interface NewNotification {
  title: string;
  message: string;
}

export const NOTIFICATION_STATUS = {
  SUCCESS: "Success",
  WARNING: "Warning",
  ERROR: "Error",
};

export function addSuccessNotification(notificationDetails: NewNotification) {
  Store.addNotification({
    title: notificationDetails.title,
    message: notificationDetails.message,
    type: "success",
    insert: "top",
    container: "top-right",
    animationIn: ["animate__animated", "animate__fadeIn"],
    animationOut: ["animate__animated", "animate__fadeOut"],
    dismiss: {
      duration: 5000,
      onScreen: true,
    },
  });
}

export function addWarningNotification(notificationDetails: NewNotification) {
  Store.addNotification({
    title: notificationDetails.title,
    message: notificationDetails.message,
    type: "warning",
    insert: "top",
    container: "top-right",
    animationIn: ["animate__animated", "animate__fadeIn"],
    animationOut: ["animate__animated", "animate__fadeOut"],
    dismiss: {
      duration: 5000,
      onScreen: true,
    },
  });
}

export function addErrorNotification(notificationDetails: NewNotification) {
  Store.addNotification({
    title: notificationDetails.title,
    message: notificationDetails.message,
    type: "success",
    insert: "top",
    container: "top-right",
    animationIn: ["animate__animated", "animate__fadeIn"],
    animationOut: ["animate__animated", "animate__fadeOut"],
    dismiss: {
      duration: 5000,
      onScreen: true,
    },
  });
}
