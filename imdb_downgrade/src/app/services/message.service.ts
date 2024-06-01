import {Injectable} from '@angular/core';
import {ToastrService} from "ngx-toastr";

@Injectable({
  providedIn: 'root'
})
export class MessageService {

  constructor(private toastService: ToastrService) {
  }

  public showSuccess(title: string, message: string): void {
    this.toastService.success(message, title);
  }

  public showInfo(title: string, message: string): void {
    this.toastService.info(message, title);
  }

  public showWarning(title: string, message: string): void {
    this.toastService.warning(message, title);
  }

  public showError(title: string, message: string): void {
    this.toastService.error(message, title);
  }
}
