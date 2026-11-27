import { Component, OnInit, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { AcademyService, Lesson, Exam, ExamQuestion } from './academy.service';

@Component({
  selector: 'app-root', standalone: true, imports: [CommonModule, FormsModule],
  template: `
  <div class="shell">
    <aside class="sidebar">
      <div class="brand"><span class="brand-mark">G</span><div><b>Grammatik</b><small>ACADEMY</small></div></div>
      <p class="nav-label">LEARNING</p>
      <button class="nav active" (click)="view='dashboard'">▦ <span>Dashboard</span></button>
      <button class="nav" (click)="view='courses'">▤ <span>Grammar courses</span></button>
      <button class="nav" (click)="startExam()">✎ <span>Exam simulator</span></button>
      <div class="side-bottom"><div class="avatar">D</div><div><b>Demo learner</b><small>Keep it up!</small></div></div>
    </aside>
    <main class="main">
      <header><div><span class="eyebrow">YOUR GERMAN JOURNEY</span><h1>{{ pageTitle }}</h1></div><div class="header-right"><span class="streak">✦ Daily practice</span><span class="avatar">D</span></div></header>
      <section *ngIf="view==='dashboard'">
        <div class="welcome"><div><span class="eyebrow light">WILLKOMMEN ZURÜCK</span><h2>Make your German<br>grammar unstoppable.</h2><p>Small lessons. Consistent practice. Real progress.</p><button class="primary" (click)="view='courses'">Continue learning <span>→</span></button></div><div class="hero-art">ä<br><small>DEUTSCH</small></div></div>
        <div class="section-head"><div><h2>Your learning path</h2><p>Six levels. One step at a time.</p></div><span class="muted">CEFR FRAMEWORK</span></div>
        <div class="levels">
          <button class="level-card" *ngFor="let l of levels" [class.selected]="selectedLevel===l.code" (click)="selectLevel(l.code)">
            <div class="level-top"><span class="level-code">{{l.code}}</span><span class="arrow">↗</span></div><b>{{l.name}}</b><small>{{l.description}}</small><div class="mini-track"><i [style.width]="levelProgress(l.code)+'%'"></i></div>
          </button>
        </div>
        <div class="section-head"><div><h2>{{selectedLevel}} grammar</h2><p>Choose a lesson and start practising.</p></div><button class="text-btn" (click)="view='courses'">All lessons →</button></div>
        <div class="lesson-grid">
          <article class="lesson-card" *ngFor="let lesson of lessons" (click)="openLesson(lesson)">
            <span class="lesson-tag">{{lesson.level}} · GRAMMAR</span><h3>{{lesson.title}}</h3><p>{{lesson.description}}</p><span class="start-link">Study lesson →</span>
          </article>
        </div>
      </section>

      <section *ngIf="view==='courses'">
        <div class="toolbar"><div><h2>Explore grammar</h2><p>Pick your level and build a strong foundation.</p></div><select [(ngModel)]="selectedLevel" (change)="loadLessons()"><option *ngFor="let l of levels" [value]="l.code">{{l.code}} — {{l.name}}</option></select></div>
        <div class="lesson-grid">
          <article class="lesson-card" *ngFor="let lesson of lessons" (click)="openLesson(lesson)">
            <span class="lesson-tag">{{lesson.level}} · LESSON {{lesson.order}}</span><h3>{{lesson.title}}</h3><p>{{lesson.description}}</p><span class="start-link">Open lesson →</span>
          </article>
          <div class="empty" *ngIf="!lessons.length">Loading lessons…</div>
        </div>
      </section>

      <section *ngIf="view==='lesson' && activeLesson as lesson" class="content-card">
        <button class="back" (click)="view='courses'">← Back to courses</button>
        <span class="lesson-tag">{{lesson.level}} · GRAMMAR LESSON</span><h2>{{lesson.title}}</h2><p class="lead">{{lesson.description}}</p>
        <h3>Grammar explained</h3><p class="explanation">{{lesson.explanation}}</p>
        <h3>Examples</h3><div class="example" *ngFor="let ex of parseExamples(lesson.examples)"><b>{{ex.de}}</b><span>{{ex.en}}</span></div>
        <div class="lesson-actions"><button class="primary" (click)="startExam()">Practise with an exam →</button></div>
      </section>

      <section *ngIf="view==='exam'">
        <div class="exam-head"><div><span class="lesson-tag">{{selectedLevel}} · KNOWLEDGE CHECK</span><h2>{{selectedLevel}} Grammar Exam</h2><p>Choose the best answer for each question.</p></div><div class="timer">◷ {{exam?.durationMinutes || 30}} min</div></div>
        <div class="question" *ngFor="let q of exam?.questions; let i=index"><span class="q-number">QUESTION {{i+1}} / {{exam?.questions?.length}}</span><h3>{{q.prompt}}</h3><label class="option" *ngFor="let option of q.options; let j=index" [class.chosen]="answers[q.id]===j"><input type="radio" [name]="'q'+q.id" [value]="j" [(ngModel)]="answers[q.id]"><span>{{option}}</span></label></div>
        <button class="primary submit" (click)="submitExam()" [disabled]="!exam || submitting">{{submitting?'Checking…':'Submit exam →'}}</button>
      </section>

      <section *ngIf="view==='result' && result" class="content-card result">
        <div class="result-icon">{{result.passed?'✓':'↗'}}</div><span class="eyebrow">EXAM COMPLETE</span><h2>{{result.passed?'Well done!':'Keep practising!'}}</h2><div class="score">{{result.percentage}}<small>%</small></div><p>You answered {{result.score}} out of {{result.total}} questions correctly.</p><span class="pass" [class.fail]="!result.passed">{{result.passed?'PASSED':'NOT PASSED'}} · PASS MARK 70%</span>
        <div class="review" *ngFor="let r of result.review; let i=index"><b>Question {{i+1}}</b><span [class.good]="r.correct" [class.bad]="!r.correct">{{r.correct?'Correct':'Review this topic'}}</span><p>{{r.explanation}}</p></div>
        <button class="primary" (click)="startExam()">Try again →</button>
      </section>
      <footer>GERMAN GRAMMAR ACADEMY <span>Learn with purpose. Fortschritt für Fortschritt.</span></footer>
    </main>
  </div>
  `
})
export class AppComponent implements OnInit {
  private api = inject(AcademyService);
  view = 'dashboard'; selectedLevel = 'A1'; levels: any[] = [];
  lessons: Lesson[] = []; activeLesson?: Lesson; exam?: Exam;
  answers: Record<number, number> = {}; result: any; submitting = false;
  get pageTitle() { return this.view==='lesson' ? 'Your lesson' : this.view==='exam' ? 'Test your knowledge' : this.view==='result' ? 'Your results' : this.view==='courses' ? 'Grammar courses' : 'Your dashboard'; }
  ngOnInit() {
    this.api.levels().subscribe(v => this.levels = v);
    this.loadLessons();
  }
  selectLevel(level: string) { this.selectedLevel=level; this.loadLessons(); }
  loadLessons() { this.api.lessons(this.selectedLevel).subscribe(v => this.lessons=v); }
  levelProgress(level: string) { return level==='A1' ? 28 : 0; }
  openLesson(lesson: Lesson) { this.activeLesson=lesson; this.view='lesson'; }
  parseExamples(raw: string): {de:string,en:string}[] { try { return JSON.parse(raw || '[]'); } catch { return []; } }
  startExam() {
    this.api.exam(this.selectedLevel).subscribe(ex => { this.exam=ex; this.answers={}; this.result=null; this.view='exam'; });
  }
  submitExam() {
    if (!this.exam) return;
    this.submitting=true;
    const answers: Record<string,number> = {};
    Object.entries(this.answers).forEach(([k,v]) => answers[k]=v);
    this.api.submit(this.selectedLevel, answers).subscribe({
      next: r => { this.result=r; this.view='result'; this.submitting=false; },
      error: () => this.submitting=false
    });
  }
}
