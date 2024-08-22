/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../../fetch.pb"

type Absent<T, K extends keyof T> = { [k in Exclude<keyof T, K>]?: undefined };
type OneOf<T> =
  | { [k in keyof T]?: undefined }
  | (
    keyof T extends infer K ?
      (K extends string & keyof T ? { [k in K]: T[K] } & Absent<T, K>
        : never)
    : never);

export enum OperationStatus {
  OPERATION_STATUS_TRY_AGAIN = "OPERATION_STATUS_TRY_AGAIN",
  OPERATION_STATUS_FAILURE = "OPERATION_STATUS_FAILURE",
  OPERATION_STATUS_SUCCESS = "OPERATION_STATUS_SUCCESS",
}

export type ValidatorFieldSet = {
  fieldsAsHex?: string[]
}

export type WithdrawalCredentialProofCollectionTStateRootProof = {
  beaconStateRoot?: string
  proof?: string
}


type BaseWithdrawalCredentialProofCollection = {
  validatorIndices?: string[]
  validatorFieldProofs?: string[]
  validatorFields?: ValidatorFieldSet[]
}

export type WithdrawalCredentialProofCollection = BaseWithdrawalCredentialProofCollection
  & OneOf<{ stateRootProof: WithdrawalCredentialProofCollectionTStateRootProof }>

export type CheckpointProofCollectionTValidatorBalancesRootProof = {
  validatorBalancesRoot?: string
  proof?: string
}

export type CheckpointProofCollectionTBalanceProof = {
  pubKeyHash?: string
  balanceRoot?: string
  proof?: string
}


type BaseCheckpointProofCollection = {
  balanceProofs?: CheckpointProofCollectionTBalanceProof[]
}

export type CheckpointProofCollection = BaseCheckpointProofCollection
  & OneOf<{ validatorBalancesRootProof: CheckpointProofCollectionTValidatorBalancesRootProof }>

export type GetCheckpointProofsRequest = {
  chainId?: string
  eigenpodAddress?: string
}

export type GetCheckpointProofsResponse = {
  status?: OperationStatus
  proofs?: CheckpointProofCollection
}

export type GetWithdrawalCredentialsProofsRequest = {
  chainId?: string
  eigenpodAddress?: string
}

export type GetWithdrawalCredentialsProofsResponse = {
  status?: OperationStatus
  oracleBeaconTimestamp?: string
  proofs?: WithdrawalCredentialProofCollection
}

export class ProofsService {
  static GetCheckpointProofs(req: GetCheckpointProofsRequest, initReq?: fm.InitReq): Promise<GetCheckpointProofsResponse> {
    return fm.fetchReq<GetCheckpointProofsRequest, GetCheckpointProofsResponse>(`/v1/checkpoint/proof`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetWithdrawalCredentialsProofs(req: GetWithdrawalCredentialsProofsRequest, initReq?: fm.InitReq): Promise<GetWithdrawalCredentialsProofsResponse> {
    return fm.fetchReq<GetWithdrawalCredentialsProofsRequest, GetWithdrawalCredentialsProofsResponse>(`/v1/withdrawal/proof`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
}