/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../../fetch.pb"
export type GenerateClaimProofRequest = {
  earnerAddress?: string
  tokens?: string[]
}

export type EarnerLeaf = {
  earner?: string
  earnerTokenRoot?: string
}

export type TokenLeaf = {
  token?: string
  cumulativeEarnings?: string
}

export type Proof = {
  root?: string
  rootIndex?: number
  earnerIndex?: number
  earnerTreeProof?: string
  earnerLeaf?: EarnerLeaf
  leafIndices?: number[]
  tokenTreeProofs?: string[]
  tokenLeaves?: TokenLeaf[]
}

export type GenerateClaimProofResponse = {
  proof?: Proof
}

export type GetEarnedTokensForEarnerRequest = {
  earnerAddress?: string
}

export type TokenAmount = {
  token?: string
  amount?: string
}

export type GetEarnedTokensForEarnerResponse = {
  tokens?: TokenAmount[]
}

export class Rewards {
  static GenerateClaimProof(req: GenerateClaimProofRequest, initReq?: fm.InitReq): Promise<GenerateClaimProofResponse> {
    return fm.fetchReq<GenerateClaimProofRequest, GenerateClaimProofResponse>(`/v1/claim-proof`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetEarnedTokensForEarner(req: GetEarnedTokensForEarnerRequest, initReq?: fm.InitReq): Promise<GetEarnedTokensForEarnerResponse> {
    return fm.fetchReq<GetEarnedTokensForEarnerRequest, GetEarnedTokensForEarnerResponse>(`/v1/earned-tokens`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
}