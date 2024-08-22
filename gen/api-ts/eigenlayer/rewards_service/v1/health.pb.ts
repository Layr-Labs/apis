/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../../fetch.pb"
import * as EigenlayerHealthHealth from "../../health/health.pb"
export class Health {
  static HealthCheck(req: EigenlayerHealthHealth.HealthCheckRequest, initReq?: fm.InitReq): Promise<EigenlayerHealthHealth.HealthCheckResponse> {
    return fm.fetchReq<EigenlayerHealthHealth.HealthCheckRequest, EigenlayerHealthHealth.HealthCheckResponse>(`/v1/health?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static ReadyCheck(req: EigenlayerHealthHealth.ReadyRequest, initReq?: fm.InitReq): Promise<EigenlayerHealthHealth.ReadyResponse> {
    return fm.fetchReq<EigenlayerHealthHealth.ReadyRequest, EigenlayerHealthHealth.ReadyResponse>(`/v1/ready?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
}