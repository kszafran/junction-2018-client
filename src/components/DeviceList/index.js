import React from "react";
import Device from "../Device";
import "./index.css";

const DeviceList = ({ devices }) => {
  return (
    <React.Fragment>
      <div className="list-title">
        <div className="list-header name">Device Name</div>
        <div className="list-header status">Status</div>
      </div>
      <div className="device-list">
        {devices.map((e, i) => (
          <Device data={e} key={i} />
        ))}
      </div>
    </React.Fragment>
  );
};

export default DeviceList;
