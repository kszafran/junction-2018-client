import React from "react";
import { Link } from "react-router-dom";
import "./index.css";

const Device = ({ data }) => {
  let { name, data_history, cisco_health_data } = data;

  let displayName = name;
  let userId = cisco_health_data.userId;
  let deviceType = cisco_health_data.deviceType;
  if (userId && deviceType) {
    displayName = `${name} (${userId} at ${deviceType})`
  } else if (userId) {
    name += `${name} (${userId})`
  } else if (deviceType) {
    name += `${name} (${deviceType})`
  }

  const currentHealthData = cisco_health_data
    ? cisco_health_data.healthScore
    : 0;

  let healthColor;
  if (currentHealthData <= 3) {
    healthColor = "red-color";
  } else if (currentHealthData <= 5) {
    healthColor = "orange-color";
  } else if (currentHealthData <= 7) {
    healthColor = "yellow-color";
  } else {
    healthColor = "green-color";
  }

  let button;
  if (data_history) {
      button =
          <Link to={`/device/${name}`}>
            <button className="detail">Details</button>
          </Link>
  }

  return (
    <div>
      <div className="device">
        <div className="name">{displayName}</div>
        <div className={`status ${healthColor}`}>{currentHealthData}</div>
        {button}
      </div>
    </div>
  );
};

export default Device;
