//
// Copyright 2019 Insolar Technologies GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by insgocc. DO NOT EDIT.
// source template in logicrunner/preprocessor/templates

package testwallet

import (
	XXX_contract "github.com/insolar/assured-ledger/ledger-core/insolar/contract"
	XXX_isolation "github.com/insolar/assured-ledger/ledger-core/insolar/contract/isolation"
	"github.com/insolar/assured-ledger/ledger-core/reference"
	"github.com/insolar/assured-ledger/ledger-core/runner/executor/common/foundation"
)

// ClassReference to class of this contract
// error checking hides in generator
var ClassReference, _ = reference.GlobalFromString("insolar:0AAABAqiF7kGalgYGa1bKDmA33RKr0lfmdtIZr73_tMU")

// Wallet holds proxy type
type Wallet struct {
	Reference   reference.Global
	Class       reference.Global
	Code        reference.Global
	ProxyHelper XXX_contract.ProxyHelper
}

// ContractConstructorHolder holds logic with object construction
type ContractConstructorHolder struct {
	proxyHelper     XXX_contract.ProxyHelper
	constructorName string
	argsSerialized  []byte
}

// AsChild saves object as child
func (r *ContractConstructorHolder) AsChild(objRef reference.Global) (*Wallet, error) {
	var ph = r.proxyHelper
	ret, err := ph.CallConstructor(objRef, ClassReference, r.constructorName, r.argsSerialized)
	if err != nil {
		return nil, err
	}

	var ref reference.Global
	var constructorError *foundation.Error
	resultContainer := foundation.Result{
		Returns: []interface{}{&ref, &constructorError},
	}
	err = ph.Deserialize(ret, &resultContainer)
	if err != nil {
		return nil, err
	}

	if resultContainer.Error != nil {
		return nil, resultContainer.Error
	}

	if constructorError != nil {
		return nil, constructorError
	}

	return &Wallet{Reference: ref}, nil
}

// GetObject returns proxy object
func GetObject(foundation foundation.ContractFoundation, ref reference.Global) *Wallet {
	if !ref.IsObjectReference() {
		return nil
	}
	return &Wallet{Reference: ref, ProxyHelper: foundation.CurrentProxyCtx()}
}

// GetClass returns reference to the class
func GetClass() reference.Global {
	return ClassReference
}

// New is constructor
func New(foundation foundation.ContractFoundation) *ContractConstructorHolder {
	var args [0]interface{}

	var argsSerialized []byte
	err := foundation.CurrentProxyCtx().Serialize(args, &argsSerialized)
	if err != nil {
		panic(err)
	}

	return &ContractConstructorHolder{
		constructorName: "New",
		argsSerialized:  argsSerialized,
		proxyHelper:     foundation.CurrentProxyCtx(),
	}
}

// GetReference returns reference of the object
func (r *Wallet) GetReference() reference.Global {
	return r.Reference
}

// GetClass returns reference to the code
func (r *Wallet) GetClass() (reference.Global, error) {
	var ph = r.ProxyHelper
	if r.Class.IsEmpty() {
		ret := [2]interface{}{}
		var ret0 reference.Global
		ret[0] = &ret0
		var ret1 *foundation.Error
		ret[1] = &ret1

		res, err := ph.CallMethod(
			r.Reference, XXX_isolation.CallIntolerable, XXX_isolation.CallValidated, false, "GetClass", make([]byte, 0), ClassReference)
		if err != nil {
			return ret0, err
		}

		err = ph.Deserialize(res, &ret)
		if err != nil {
			return ret0, err
		}

		if ret1 != nil {
			return ret0, ret1
		}

		r.Class = ret0
	}

	return r.Class, nil

}

// GetCode returns reference to the code
func (r *Wallet) GetCode() (reference.Global, error) {
	var ph = r.ProxyHelper
	if r.Code.IsEmpty() {
		ret := [2]interface{}{}
		var ret0 reference.Global
		ret[0] = &ret0
		var ret1 *foundation.Error
		ret[1] = &ret1

		res, err := ph.CallMethod(
			r.Reference, XXX_isolation.CallIntolerable, XXX_isolation.CallValidated, false, "GetCode", make([]byte, 0), ClassReference)
		if err != nil {
			return ret0, err
		}

		err = ph.Deserialize(res, &ret)
		if err != nil {
			return ret0, err
		}

		if ret1 != nil {
			return ret0, ret1
		}

		r.Code = ret0
	}

	return r.Code, nil
}

// GetBalance is proxy generated method
func (r *Wallet) GetBalanceAsMutable() (uint32, error) {
	var args [0]interface{}

	var argsSerialized []byte

	ret := make([]interface{}, 2)
	var ret0 uint32
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	var ph = r.ProxyHelper

	err := ph.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := ph.CallMethod(r.Reference, XXX_isolation.CallTolerable, XXX_isolation.CallDirty, false, "GetBalance", argsSerialized, ClassReference)
	if err != nil {
		return ret0, err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = ph.Deserialize(res, &resultContainer)
	if err != nil {
		return ret0, err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return ret0, err
	}
	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// GetBalanceAsImmutable is proxy generated method
func (r *Wallet) GetBalance() (uint32, error) {
	var args [0]interface{}

	var argsSerialized []byte

	ret := make([]interface{}, 2)
	var ret0 uint32
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	var ph = r.ProxyHelper

	err := ph.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := ph.CallMethod(
		r.Reference, XXX_isolation.CallIntolerable, XXX_isolation.CallValidated, false, "GetBalance", argsSerialized, ClassReference)
	if err != nil {
		return ret0, err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = ph.Deserialize(res, &resultContainer)
	if err != nil {
		return ret0, err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return ret0, err
	}
	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// Accept is proxy generated method
func (r *Wallet) Accept(amount uint32) error {
	var args [1]interface{}
	args[0] = amount

	var argsSerialized []byte

	ret := make([]interface{}, 1)
	var ret0 *foundation.Error
	ret[0] = &ret0

	var ph = r.ProxyHelper

	err := ph.Serialize(args, &argsSerialized)
	if err != nil {
		return err
	}

	res, err := ph.CallMethod(r.Reference, XXX_isolation.CallTolerable, XXX_isolation.CallDirty, false, "Accept", argsSerialized, ClassReference)
	if err != nil {
		return err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = ph.Deserialize(res, &resultContainer)
	if err != nil {
		return err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return err
	}
	if ret0 != nil {
		return ret0
	}
	return nil
}

// AcceptAsImmutable is proxy generated method
func (r *Wallet) AcceptAsImmutable(amount uint32) error {
	var args [1]interface{}
	args[0] = amount

	var argsSerialized []byte

	ret := make([]interface{}, 1)
	var ret0 *foundation.Error
	ret[0] = &ret0

	var ph = r.ProxyHelper

	err := ph.Serialize(args, &argsSerialized)
	if err != nil {
		return err
	}

	res, err := ph.CallMethod(
		r.Reference, XXX_isolation.CallIntolerable, XXX_isolation.CallValidated, false, "Accept", argsSerialized, ClassReference)
	if err != nil {
		return err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = ph.Deserialize(res, &resultContainer)
	if err != nil {
		return err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return err
	}
	if ret0 != nil {
		return ret0
	}
	return nil
}

// Transfer is proxy generated method
func (r *Wallet) Transfer(toWallet reference.Global, amount uint32) error {
	var args [2]interface{}
	args[0] = toWallet
	args[1] = amount

	var argsSerialized []byte

	ret := make([]interface{}, 1)
	var ret0 *foundation.Error
	ret[0] = &ret0

	var ph = r.ProxyHelper

	err := ph.Serialize(args, &argsSerialized)
	if err != nil {
		return err
	}

	res, err := ph.CallMethod(r.Reference, XXX_isolation.CallTolerable, XXX_isolation.CallDirty, false, "Transfer", argsSerialized, ClassReference)
	if err != nil {
		return err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = ph.Deserialize(res, &resultContainer)
	if err != nil {
		return err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return err
	}
	if ret0 != nil {
		return ret0
	}
	return nil
}

// TransferAsImmutable is proxy generated method
func (r *Wallet) TransferAsImmutable(toWallet reference.Global, amount uint32) error {
	var args [2]interface{}
	args[0] = toWallet
	args[1] = amount

	var argsSerialized []byte

	ret := make([]interface{}, 1)
	var ret0 *foundation.Error
	ret[0] = &ret0

	var ph = r.ProxyHelper

	err := ph.Serialize(args, &argsSerialized)
	if err != nil {
		return err
	}

	res, err := ph.CallMethod(
		r.Reference, XXX_isolation.CallIntolerable, XXX_isolation.CallValidated, false, "Transfer", argsSerialized, ClassReference)
	if err != nil {
		return err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = ph.Deserialize(res, &resultContainer)
	if err != nil {
		return err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return err
	}
	if ret0 != nil {
		return ret0
	}
	return nil
}

// Destroy is proxy generated method
func (r *Wallet) Destroy() error {
	var args [0]interface{}

	var argsSerialized []byte

	ret := make([]interface{}, 1)
	var ret0 *foundation.Error
	ret[0] = &ret0

	var ph = r.ProxyHelper

	err := ph.Serialize(args, &argsSerialized)
	if err != nil {
		return err
	}

	res, err := ph.CallMethod(r.Reference, XXX_isolation.CallTolerable, XXX_isolation.CallDirty, false, "Destroy", argsSerialized, ClassReference)
	if err != nil {
		return err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = ph.Deserialize(res, &resultContainer)
	if err != nil {
		return err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return err
	}
	if ret0 != nil {
		return ret0
	}
	return nil
}

// DestroyAsImmutable is proxy generated method
func (r *Wallet) DestroyAsImmutable() error {
	var args [0]interface{}

	var argsSerialized []byte

	ret := make([]interface{}, 1)
	var ret0 *foundation.Error
	ret[0] = &ret0

	var ph = r.ProxyHelper

	err := ph.Serialize(args, &argsSerialized)
	if err != nil {
		return err
	}

	res, err := ph.CallMethod(
		r.Reference, XXX_isolation.CallIntolerable, XXX_isolation.CallValidated, false, "Destroy", argsSerialized, ClassReference)
	if err != nil {
		return err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = ph.Deserialize(res, &resultContainer)
	if err != nil {
		return err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return err
	}
	if ret0 != nil {
		return ret0
	}
	return nil
}
