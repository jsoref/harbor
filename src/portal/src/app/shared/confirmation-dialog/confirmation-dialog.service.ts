// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
import { Injectable } from '@angular/core';
import { Subject } from "rxjs";

import { ConfirmationMessage } from './confirmation-message';
import { ConfirmationAcknowledgement } from './confirmation-state-message';

@Injectable()
export class ConfirmationDialogService {
    confirmationAnnouncedSource = new Subject<ConfirmationMessage>();
    confirmationConfirmSource = new Subject<ConfirmationAcknowledgement>();

    confirmationAnnounced$ = this.confirmationAnnouncedSource.asObservable();
    confirmationConfirm$ = this.confirmationConfirmSource.asObservable();

    // User confirm the action
    public confirm(ack: ConfirmationAcknowledgement): void {
        this.confirmationConfirmSource.next(ack);
    }

    // User cancel the action
    public cancel(ack: ConfirmationAcknowledgement): void {
        this.confirm(ack);
    }

    // Open the confirmation dialog
    public openComfirmDialog(message: ConfirmationMessage): void {
        this.confirmationAnnouncedSource.next(message);
    }
}
