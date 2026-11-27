import { Injectable, inject } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

export interface Lesson {
  id: number; level: string; title: string; slug: string;
  description: string; explanation: string; examples: string; order: number;
}
export interface ExamQuestion { id: number; prompt: string; options: string[]; }
export interface Exam { level: string; durationMinutes: number; passMark: number; questions: ExamQuestion[]; }

@Injectable({ providedIn: 'root' })
export class AcademyService {
  private http = inject(HttpClient);
  levels(): Observable<any[]> { return this.http.get<any[]>('/api/levels'); }
  lessons(level: string): Observable<Lesson[]> { return this.http.get<Lesson[]>(`/api/lessons?level=${level}`); }
  exam(level: string): Observable<Exam> { return this.http.get<Exam>(`/api/exams/${level}`); }
  submit(level: string, answers: Record<string, number>) {
    return this.http.post<any>(`/api/exams/${level}/submit`, { userId: 'demo', answers });
  }
  progress() { return this.http.get<any>('/api/progress?userId=demo'); }
}
