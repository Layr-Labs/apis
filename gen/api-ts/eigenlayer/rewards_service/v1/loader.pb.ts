/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../../fetch.pb"
export type ReloadProofsRequest = {
}

export type ReloadProofsResponse = {
}

export class ProofsLoaderService {
  static ReloadProofs(req: ReloadProofsRequest, initReq?: fm.InitReq): Promise<ReloadProofsResponse> {
    return fm.fetchReq<ReloadProofsRequest, ReloadProofsResponse>(`/eigenlayer.apis.loader.v1.ProofsLoaderService/ReloadProofs`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
}